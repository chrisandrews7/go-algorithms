package main

import (
	"testing"
)

func TestLLInsertAfter(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)

	list.InsertAfter(&node1)
	list.InsertAfter(&node2)

	if list.head != &node1 {
		t.Error("Head was not set")
	}
	if list.tail != &node2 {
		t.Error("Tail was not set")
	}
	if size := list.length; size != 2 {
		t.Errorf("Wrong length, expected 2 and got %d", size)
	}
	if node1.next != &node2 {
		t.Error("Nodes are was not linked")
	}
}

func TestLLInsertBefore(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)

	list.InsertBefore(&node1)
	list.InsertBefore(&node2)

	if list.head != &node2 {
		t.Error("Head was not set")
	}
	if list.tail != &node1 {
		t.Error("Tail was not set")
	}
	if size := list.length; size != 2 {
		t.Errorf("Wrong length, expected 2 and got %d", size)
	}
	if node2.next != &node1 {
		t.Error("Nodes are was not linked")
	}
}

func TestLLRemoveFirst(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)

	list.InsertAfter(&node1)
	list.InsertAfter(&node2)

	popped := list.RemoveFirst()

	if node1.next == &node2 || list.head != &node2 || list.tail != &node2 {
		t.Error("Node 1 was not removed")
	}
	if popped != &node1 {
		t.Error("Node 1 was not returned")
	}
	if size := list.length; size != 1 {
		t.Errorf("Wrong length, expected 1 and got %d", size)
	}
}

func TestLLRemoveFirstWithOneNode(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)

	list.InsertAfter(&node1)

	popped := list.RemoveFirst()

	if list.head != nil || list.tail != nil {
		t.Error("Node was not removed")
	}
	if popped != &node1 {
		t.Error("Node was not returned")
	}
	if size := list.length; size != 0 {
		t.Errorf("Wrong length, expected 0 and got %d", size)
	}
}

func TestLLRemoveLast(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	list.InsertAfter(&node1)
	list.InsertAfter(&node2)
	list.InsertAfter(&node3)

	popped := list.RemoveLast()

	if node2.next == &node3 || list.tail != &node2 {
		t.Error("Node 3 was not removed")
	}
	if popped != &node3 {
		t.Error("Node 3 was not returned")
	}
	if size := list.length; size != 2 {
		t.Errorf("Wrong length, expected 2 and got %d", size)
	}
}

func TestLLRemoveLastWithOneNode(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)

	list.InsertAfter(&node1)

	popped := list.RemoveLast()

	if list.head != nil || list.tail != nil {
		t.Error("Node was not removed")
	}
	if popped != &node1 {
		t.Error("Node was not returned")
	}
	if size := list.length; size != 0 {
		t.Errorf("Wrong length, expected 0 and got %d", size)
	}
}

func TestLLRemove(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	list.InsertAfter(&node1)
	list.InsertAfter(&node2)
	list.InsertAfter(&node3)

	list.Remove(&node2)

	if node1.next == &node2 {
		t.Error("Node 2 was not removed")
	}
	if size := list.length; size != 2 {
		t.Errorf("Wrong length, expected 2 and got %d", size)
	}
}

func TestLLRemoveWithOneNode(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewNode(1)

	list.InsertAfter(&node1)

	list.Remove(&node1)

	if list.head != nil || list.tail != nil {
		t.Error("Node was not removed")
	}
	if size := list.length; size != 0 {
		t.Errorf("Wrong length, expected 0 and got %d", size)
	}
}

func TestLLRemoveInvalidNode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a not found panic")
		}
	}()

	list := NewSinglyLinkedList()
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	list.InsertAfter(&node1)
	list.InsertAfter(&node2)

	// node3 was never added
	list.Remove(&node3)
}
