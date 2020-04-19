package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (tree *BinarySearchTree) Insert(value int) {
	node := &TreeNode{
		value: value,
	}

	if tree.root == nil {
		tree.root = node
		return
	}

	currentNode := tree.root
	for {
		if currentNode.value < value {
			if currentNode.right == nil {
				currentNode.right = node
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = node
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinarySearchTree) Find(value int) *TreeNode {
	currentNode := tree.root
	for {
		if currentNode == nil || currentNode.value == value {
			return currentNode
		}

		if currentNode.value < value {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{}
}
