package csmt

import (
	"github.com/phoreproject/synapse/chainhash"
)


func combineHashes(left *chainhash.Hash, right *chainhash.Hash) chainhash.Hash {
	return chainhash.HashH(append(left[:], right[:]...))
}

var emptyHash = chainhash.Hash{}
var emptyTrees [256]chainhash.Hash

// EmptyTree is the hash of an empty tree.
var EmptyTree = chainhash.Hash{}

func init() {
	emptyTrees[0] = emptyHash
	for i := range emptyTrees[1:] {
		emptyTrees[i+1] = combineHashes(&emptyTrees[i], &emptyTrees[i])
	}

	EmptyTree = emptyTrees[255]
}

// isRight checks if the key is in the left or right subtree at a certain level. Level 255 is the root level.
func isRight(key chainhash.Hash, level uint8) bool {
	return key[level/8]&(1<<uint(level%8)) != 0
}

// calculateSubtreeHashWithOneLeaf calculates the hash of a subtree with only a single leaf at a certain height.
// atLevel is the height to calculate at.
func calculateSubtreeHashWithOneLeaf(key *chainhash.Hash, value *chainhash.Hash, atLevel uint8) chainhash.Hash {
	h := *value

	for i := uint8(0); i < atLevel; i++ {
		right := isRight(*key, i+1)

		// the key is in the right subtree
		if right {
			h = combineHashes(&emptyTrees[i], &h)
		} else {
			h = combineHashes(&h, &emptyTrees[i])
		}
	}

	return h
}

func insertIntoTree(t TreeDatabase, root *Node, key chainhash.Hash, value chainhash.Hash, level uint8) (*Node, error) {
	right := isRight(key, level)

	if level == 0 {
		if root != nil && !root.Empty() {
			// remove the old node if it exists
			err := t.DeleteNode(root.GetHash())
			if err != nil {
				return nil, err
			}
		}

		// bottom leafs should have no siblings and a value
		return t.NewNode(nil, nil, value)
	}

	// if this tree is empty and we're inserting, we know it's the only key in the subtree, so let's mark it as such and
	// fill in the necessary values
	if root == nil || root.Empty() {
		return t.NewSingleNode(key, value, calculateSubtreeHashWithOneLeaf(&key, &value, level))
	}

	leftHash := root.Left()
	rightHash := root.Right()

	var newLeftBranch *Node
	var newRightBranch *Node

	if leftHash != nil {
		newLeftBranch, _ = t.GetNode(*leftHash)
	}

	if rightHash != nil {
		newRightBranch, _ = t.GetNode(*rightHash)
	}

	// if there is only one key in this subtree,
	if root.IsSingle() {
		rootKey := root.GetSingleKey()

		// this operation is an update
		if rootKey.IsEqual(&key) {
			// delete the old root
			err := t.DeleteNode(root.GetHash())
			if err != nil {
				return nil, err
			}

			// calculate the new root hash for this subtree
			return t.NewSingleNode(key, value, calculateSubtreeHashWithOneLeaf(&key, &value, level))
		}

		// check if the old key goes in the left or right
		subRight := isRight(rootKey, level)

		// we know this is a single, so the left and right should be nil
		if subRight {
			rightBranchInserted, err := insertIntoTree(t, newRightBranch, rootKey, root.GetSingleValue(), level-1)
			if err != nil {
				return nil, err
			}
			newRightBranch = rightBranchInserted
		} else {
			leftBranchInserted, err := insertIntoTree(t, newLeftBranch, rootKey, root.GetSingleValue(), level-1)
			if err != nil {
				return nil, err
			}
			newLeftBranch = leftBranchInserted
		}
	}

	if right {
		rightBranchInserted, err := insertIntoTree(t, newRightBranch, key, value, level-1)
		if err != nil {
			return nil, err
		}
		newRightBranch = rightBranchInserted
	} else {
		leftBranchInserted, err := insertIntoTree(t, newLeftBranch, key, value, level-1)
		if err != nil {
			return nil, err
		}
		newLeftBranch = leftBranchInserted
	}

	lv := emptyTrees[level-1]
	if newLeftBranch != nil && !newLeftBranch.Empty() {
		lv = newLeftBranch.GetHash()
	}

	rv := emptyTrees[level-1]
	if newRightBranch != nil && !newRightBranch.Empty() {
		rv = newRightBranch.GetHash()
	}

	newHash := combineHashes(&lv, &rv)

	return t.NewNode(newLeftBranch, newRightBranch, newHash)
}