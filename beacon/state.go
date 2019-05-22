package beacon

import (
	"errors"
	"fmt"
	"time"

	"github.com/prysmaticlabs/prysm/shared/ssz"

	"github.com/phoreproject/synapse/primitives"
	"github.com/phoreproject/synapse/utils"
	logger "github.com/sirupsen/logrus"
)

// StoreBlock adds a block header to the current chain. The block should already
// have been validated by this point.
func (b *Blockchain) StoreBlock(block *primitives.Block) error {
	err := b.DB.SetBlock(*block)
	if err != nil {
		return err
	}

	return nil
}

// AddBlockToStateMap calculates the state after applying block and adds it
// to the state map.
func (b *Blockchain) AddBlockToStateMap(block *primitives.Block, verifySignature bool) ([]primitives.Receipt, *primitives.State, error) {
	return b.stateManager.AddBlockToStateMap(block, verifySignature)
}

// ProcessBlock is called when a block is received from a peer.
func (b *Blockchain) ProcessBlock(block *primitives.Block, checkTime bool, verifySignature bool) ([]primitives.Receipt, *primitives.State, error) {
	genesisTime := b.stateManager.GetGenesisTime()

	validationStart := time.Now()

	// VALIDATE BLOCK HERE
	if checkTime && (block.BlockHeader.SlotNumber*uint64(b.config.SlotDuration)+genesisTime > uint64(utils.Now().Unix()) || block.BlockHeader.SlotNumber == 0) {
		return nil, nil, errors.New("block slot too soon")
	}

	seen := b.View.Index.Has(block.BlockHeader.ParentRoot)
	if !seen {
		return nil, nil, errors.New("do not have parent block")
	}

	blockHash, err := ssz.TreeHash(block)
	if err != nil {
		return nil, nil, err
	}

	seen = b.View.Index.Has(blockHash)

	if seen {
		// we've already processed this block
		return nil, nil, nil
	}

	validationTime := time.Since(validationStart)

	blockHashStr := fmt.Sprintf("%x", blockHash)

	logger.WithFields(logger.Fields{
		"hash": blockHashStr,
		"slot": block.BlockHeader.SlotNumber,
	}).Info("processing new block")

	stateCalculationStart := time.Now()

	initialState, found := b.stateManager.GetStateForHash(block.BlockHeader.ParentRoot)
	if !found {
		return nil, nil, errors.New("could not find state for parent block")
	}

	initialJustifiedSlot := initialState.JustifiedSlot
	initialFinalizedSlot := initialState.FinalizedSlot

	receipts, newState, err := b.AddBlockToStateMap(block, verifySignature)
	if err != nil {
		return nil, nil, err
	}

	stateCalculationTime := time.Since(stateCalculationStart)

	blockStorageStart := time.Now()

	err = b.StoreBlock(block)
	if err != nil {
		return nil, nil, err
	}

	blockStorageTime := time.Since(blockStorageStart)

	//logger.Debug("applied with new state")

	node, err := b.View.Index.AddBlockNodeToIndex(block, blockHash)
	if err != nil {
		return nil, nil, err
	}

	databaseTipUpdateStart := time.Now()

	err = b.DB.TransactionalUpdate(func(transaction interface{}) error {
		// set the block node in the database
		err = b.DB.SetBlockNode(blockNodeToDisk(*node), transaction)
		if err != nil {
			return err
		}

		// update the parent node in the database
		err = b.DB.SetBlockNode(blockNodeToDisk(*node.Parent), transaction)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	databaseTipUpdateTime := time.Since(databaseTipUpdateStart)

	attestationUpdateStart := time.Now()

	for _, a := range block.BlockBody.Attestations {
		participants, err := newState.GetAttestationParticipants(a.Data, a.ParticipationBitfield, b.config, newState.Slot-1)
		if err != nil {
			return nil, nil, err
		}

		err = b.DB.SetLatestAttestationsIfNeeded(participants, a)
		if err != nil {
			return nil, nil, err
		}
	}

	attestationUpdateEnd := time.Since(attestationUpdateStart)

	//logger.Debug("updating chain head")

	updateChainHeadStart := time.Now()

	err = b.UpdateChainHead()
	if err != nil {
		return nil, nil, err
	}

	updateChainHeadTime := time.Since(updateChainHeadStart)

	connectBlockSignalStart := time.Now()

	for _, n := range b.Notifees {
		go n.ConnectBlock(block)
	}

	connectBlockSignalTime := time.Since(connectBlockSignalStart)

	finalizedStateUpdateStart := time.Now()

	finalizedNode := node.GetAncestorAtSlot(newState.FinalizedSlot)
	if finalizedNode == nil {
		return nil, nil, errors.New("could not find finalized node in block index")
	}
	finalizedState, found := b.stateManager.GetStateForHash(finalizedNode.Hash)
	if !found {
		return nil, nil, errors.New("could not find finalized block Hash in state map")
	}

	if initialFinalizedSlot != newState.FinalizedSlot {
		err := b.DB.SetFinalizedState(*finalizedState)
		if err != nil {
			return nil, nil, err
		}
	}

	finalizedNodeAndState := blockNodeAndState{finalizedNode, *finalizedState}
	b.View.finalizedHead = finalizedNodeAndState

	err = b.DB.SetFinalizedHead(finalizedNode.Hash)
	if err != nil {
		return nil, nil, err
	}

	justifiedNode := node.GetAncestorAtSlot(newState.JustifiedSlot)
	if justifiedNode == nil {
		return nil, nil, errors.New("could not find justified node in block index")
	}

	justifiedState, found := b.stateManager.GetStateForHash(justifiedNode.Hash)
	if !found {
		return nil, nil, errors.New("could not find justified block Hash in state map")
	}
	justifiedNodeAndState := blockNodeAndState{justifiedNode, *justifiedState}
	b.View.justifiedHead = justifiedNodeAndState

	if initialJustifiedSlot != newState.JustifiedSlot {
		err := b.DB.SetJustifiedState(*justifiedState)
		if err != nil {
			return nil, nil, err
		}
	}

	err = b.DB.SetJustifiedHead(justifiedNode.Hash)
	if err != nil {
		return nil, nil, err
	}

	finalizedStateUpdateTime := time.Since(finalizedStateUpdateStart)

	stateCleanupStart := time.Now()

	err = b.stateManager.DeleteStateBeforeFinalizedSlot(finalizedNode.Slot)
	if err != nil {
		return nil, nil, err
	}

	stateCleanupTime := time.Since(stateCleanupStart)

	logger.WithFields(logger.Fields{
		"validation":         validationTime,
		"stateCalculation":   stateCalculationTime,
		"storage":            blockStorageTime,
		"databaseTipUpdate":  databaseTipUpdateTime,
		"attestationUpdate":  attestationUpdateEnd,
		"updateChainHead":    updateChainHeadTime,
		"connectBlockSignal": connectBlockSignalTime,
		"finalizedUpdate":    finalizedStateUpdateTime,
		"stateCleanup":       stateCleanupTime,
		"totalTime":          time.Since(validationStart),
	})

	return receipts, newState, nil
}

// GetState gets a copy of the current state of the blockchain.
func (b *Blockchain) GetState() primitives.State {
	return b.stateManager.GetHeadState()
}
