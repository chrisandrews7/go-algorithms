package main

import (
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	isLess := func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}
	heap := NewMaxBinaryHeap(isLess)

	heap.Enqueue(41)
	heap.Enqueue(39)
	heap.Enqueue(33)
	heap.Enqueue(18)
	heap.Enqueue(27)
	heap.Enqueue(12)
	heap.Enqueue(55)

	//         55
	//    39        41
	// 18    27  12    33

	if removed := heap.Dequeue(); removed != 55 {
		t.Errorf("Expected 55, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != 41 {
		t.Errorf("Expected 41, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != 39 {
		t.Errorf("Expected 39, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != 33 {
		t.Errorf("Expected 33, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != 27 {
		t.Errorf("Expected 27, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != 18 {
		t.Errorf("Expected 18, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != 12 {
		t.Errorf("Expected 12, got %d", removed)
	}
	if removed := heap.Dequeue(); removed != nil {
		t.Errorf("Expected nil, got %d", removed)
	}
}
