package main

import (
	"testing"
)

func TestBSTInsert(t *testing.T) {
	tree := NewBinarySearchTree()

	node1 := NewTreeNode(30)
	node2 := NewTreeNode(40)
	node3 := NewTreeNode(10)
	node4 := NewTreeNode(55)

	tree.Insert(&node1)
	tree.Insert(&node2)
	tree.Insert(&node3)
	tree.Insert(&node4)

	//      30
	//   10    40
	//            55

	if tree.root != &node1 {
		t.Error("Tree root is not node1")
	}
	if tree.root.right != &node2 {
		t.Error("First right branch should be 40")
	}
	if tree.root.left != &node3 {
		t.Error("First left branch should be 10")
	}
	if tree.root.right.right != &node4 {
		t.Error("Second right branch should be 55")
	}
}

func TestBSTSearch(t *testing.T) {
	tree := NewBinarySearchTree()

	node1 := NewTreeNode(30)
	node2 := NewTreeNode(40)
	node3 := NewTreeNode(10)
	node4 := NewTreeNode(55)
	node5 := NewTreeNode(60)
	node6 := NewTreeNode(45)

	tree.Insert(&node1)
	tree.Insert(&node2)
	tree.Insert(&node3)
	tree.Insert(&node4)
	tree.Insert(&node5)
	tree.Insert(&node6)

	//      30
	//   10    40
	//            55
	//         45    60

	result := tree.Find(45)

	if result != &node6 {
		t.Error("45 wasn't found")
	}
}

func TestBSTInvalidSearch(t *testing.T) {
	tree := NewBinarySearchTree()

	node1 := NewTreeNode(30)
	node2 := NewTreeNode(40)
	tree.Insert(&node1)
	tree.Insert(&node2)

	result := tree.Find(100)

	if result != nil {
		t.Error("100 should not be found")
	}
}
