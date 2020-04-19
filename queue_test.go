package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	node1 := 1
	node2 := 2

	queue.Insert(node1)
	queue.Insert(node2)

	if queue.Remove() != node1 {
		t.Error("Expected node1 out first")
	}
	if queue.Remove() != node2 {
		t.Error("Expected node2 out last")
	}
}
