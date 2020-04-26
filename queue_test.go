package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	node1 := 1
	node2 := 2

	queue.Enqueue(node1)
	queue.Enqueue(node2)

	if queue.Dequeue() != node1 {
		t.Error("Expected node1 out first")
	}
	if queue.Dequeue() != node2 {
		t.Error("Expected node2 out last")
	}
}
