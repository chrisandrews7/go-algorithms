package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	node1 := NewNode(1)
	node2 := NewNode(2)

	stack.Insert(&node1)
	stack.Insert(&node2)

	popped1 := stack.Remove()
	popped2 := stack.Remove()

	if popped1 != &node2 {
		t.Error("Expected node2 out first")
	}
	if popped2 != &node1 {
		t.Error("Expected node1 out last")
	}
}
