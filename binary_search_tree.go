package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (tree *BinarySearchTree) Insert(node *TreeNode) {
	var insert func(currentNode *TreeNode)
	insert = func(currentNode *TreeNode) {
		if currentNode.value < node.value {
			if currentNode.right == nil {
				currentNode.right = node
				return
			}
			insert(currentNode.right)
		} else {
			if currentNode.left == nil {
				currentNode.left = node
				return
			}
			insert(currentNode.left)
		}
	}

	if tree.root == nil {
		tree.root = node
		return
	}

	insert(tree.root)
}

func (tree *BinarySearchTree) Find(value int) *TreeNode {
	var find func(currentNode *TreeNode) *TreeNode
	find = func(currentNode *TreeNode) *TreeNode {
		if currentNode == nil || currentNode.value == value {
			return currentNode
		}

		if currentNode.value < value {
			return find(currentNode.right)
		} else {
			return find(currentNode.left)
		}
	}

	return find(tree.root)
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{}
}

func NewTreeNode(value int) TreeNode {
	return TreeNode{
		value: value,
	}
}
