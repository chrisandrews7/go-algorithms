package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	node1 := 1
	node2 := 2

	stack.Insert(node1)
	stack.Insert(node2)

	if stack.Remove() != node2 {
		t.Error("Expected node2 out first")
	}
	if stack.Remove() != node1 {
		t.Error("Expected node1 out last")
	}
}
