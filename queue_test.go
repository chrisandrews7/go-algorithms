package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	node1 := NewNode(1)
	node2 := NewNode(2)

	queue.Insert(&node1)
	queue.Insert(&node2)

	popped1 := queue.Remove()
	popped2 := queue.Remove()

	if popped1 != &node1 {
		t.Error("Expected node1 out first")
	}
	if popped2 != &node2 {
		t.Error("Expected node2 out last")
	}
}
