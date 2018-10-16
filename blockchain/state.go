package blockchain

import (
	"errors"
	"math"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/phoreproject/synapse/primitives"
	"github.com/phoreproject/synapse/transaction"
)

// ValidateAttestation checks attestation invariants and the BLS signature.
func (b Blockchain) ValidateAttestation(attestation transaction.Attestation, block primitives.BlockHeader, parentBlock primitives.BlockHeader, c Config) error {
	if attestation.Slot > parentBlock.SlotNumber {
		return errors.New("attestation slot number too high")
	}

	if !(attestation.Slot >= uint64(math.Max(float64(parentBlock.SlotNumber-uint64(c.CycleLength)+1), 0))) {
		return errors.New("attestation slot number too low")
	}

	if attestation.JustifiedSlot > b.state.Crystallized.LastJustifiedSlot {
		return errors.New("last justified slot should be less than or equal to the crystallized slot")
	}

	justifiedBlock, err := b.index.GetBlockNodeByHash(attestation.JustifiedBlockHash)
	if err != nil {
		return errors.New("justified block not in index")
	}

	if justifiedBlock.SlotNumber != attestation.Slot {
		return errors.New("justified slot does not match attestation")
	}

	// if (!len(attestation.att))

	// TODO: validate BLS sig

	return nil
}

// AddBlock adds a block header to the current chain. The block should already
// have been validated by this point.
func (b *Blockchain) AddBlock(h primitives.BlockHeader) error {
	var parent *BlockNode
	if !(h.AncestorHashes[0] == zeroHash && len(b.chain) == 0) {
		p, err := b.index.GetBlockNodeByHash(h.AncestorHashes[0])
		parent = p
		if err != nil {
			return err
		}
	}

	height := uint64(0)
	if parent != nil {
		height = parent.Height + 1
	}

	node := &BlockNode{BlockHeader: h, PrevNode: parent, Height: height}

	// Add block to the index
	b.index.AddNode(node)

	b.UpdateChainHead(node)

	return nil
}

// ProcessBlock is called when a block is received from a peer.
func (b Blockchain) ProcessBlock(block primitives.Block) error {
	err := b.ValidateIncomingBlock(block)
	if err != nil {
		return err
	}

	b.AddBlock(block.BlockHeader)

	return nil
}

// GetNewRecentBlockHashes will take a list of recent block hashes and
// shift them to the right, filling them in with the parentHash provided.
func GetNewRecentBlockHashes(oldHashes []*chainhash.Hash, parentSlot uint32, currentSlot uint32, parentHash *chainhash.Hash) []*chainhash.Hash {
	d := currentSlot - parentSlot
	newHashes := oldHashes[d:]
	numberToAdd := int(d)
	if numberToAdd > len(oldHashes) {
		numberToAdd = len(oldHashes)
	}
	for i := 0; i < numberToAdd; i++ {
		newHashes = append(newHashes, parentHash)
	}
	return newHashes
}

// UpdateAncestorHashes fills in the parent hash in ancestor hashes
// where the ith element represents the 2**i past block.
func UpdateAncestorHashes(parentAncestorHashes []chainhash.Hash, parentSlotNumber uint64, parentHash chainhash.Hash) []chainhash.Hash {
	newAncestorHashes := parentAncestorHashes[:]
	for i := uint(0); i < 32; i++ {
		if parentSlotNumber%(1<<i) == 0 {
			newAncestorHashes[i] = parentHash
		}
	}
	return newAncestorHashes
}

// GetShardsAndCommitteesForSlot gets the committee for each shard.
func (b Blockchain) GetShardsAndCommitteesForSlot(slot uint64) []primitives.ShardAndCommittee {
	earliestSlotInArray := b.state.Crystallized.LastStateRecalculation - uint64(b.config.CycleLength)
	return b.state.Crystallized.ShardAndCommitteeForSlots[slot-earliestSlotInArray]
}

// ValidateIncomingBlock runs a couple of checks on an incoming block.
func (b Blockchain) ValidateIncomingBlock(newBlock primitives.Block) error {
	if len(newBlock.AncestorHashes) != 32 {
		return errors.New("ancestorHashes improperly formed")
	}

	parentBlock, err := b.index.GetBlockNodeByHash(newBlock.AncestorHashes[0])
	if err != nil {
		return err
	}

	newHashes := UpdateAncestorHashes(parentBlock.AncestorHashes, parentBlock.SlotNumber, parentBlock.Hash())
	for i := range newBlock.AncestorHashes {
		if newHashes[i] != newBlock.AncestorHashes[i] {
			return errors.New("ancestor hashes don't match expected value")
		}
	}

	for _, tx := range newBlock.Transactions {
		if sat, success := tx.Data.(transaction.SubmitAttestationTransaction); success {
			err := b.ValidateAttestation(sat.Attestation, newBlock.BlockHeader, parentBlock.BlockHeader, b.config)
			if err != nil {
				return err
			}
			b.state.Active.PendingAttestations = append(b.state.Active.PendingAttestations, sat.Attestation)
		}
		if _, success := tx.Data.(transaction.LoginTransaction); success {
			b.state.Active.PendingActions = append(b.state.Active.PendingActions, tx)
		}
		if _, success := tx.Data.(transaction.LogoutTransaction); success {
			b.state.Active.PendingActions = append(b.state.Active.PendingActions, tx)
		}
		if _, success := tx.Data.(transaction.RegisterTransaction); success {
			b.state.Active.PendingActions = append(b.state.Active.PendingActions, tx)
		}
	}

	// TODO: check attestation from proposer
	// TODO: check local time
	return nil
}
