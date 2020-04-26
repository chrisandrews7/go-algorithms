package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	node1 := 1
	node2 := 2

	stack.Enqueue(node1)
	stack.Enqueue(node2)

	if stack.Dequeue() != node2 {
		t.Error("Expected node2 out first")
	}
	if stack.Dequeue() != node1 {
		t.Error("Expected node1 out last")
	}
}
