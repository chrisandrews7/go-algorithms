package main

import (
	"reflect"
	"testing"
)

func TestBHInsert(t *testing.T) {
	heap := NewMaxBinaryHeap()

	heap.Insert(41)
	heap.Insert(39)
	heap.Insert(33)
	heap.Insert(18)
	heap.Insert(27)
	heap.Insert(12)
	heap.Insert(55)

	expectedHeap := []int{55, 39, 41, 18, 27, 12, 33}
	//         55
	//    39        41
	// 18    27  12    33

	if !reflect.DeepEqual(heap.values, expectedHeap) {
		t.Errorf("Heap %d is not in the correct order", heap.values)
	}
}

func TestBHExtractMax(t *testing.T) {
	heap := NewMaxBinaryHeap()

	heap.Insert(41)
	heap.Insert(39)
	heap.Insert(33)
	heap.Insert(18)
	heap.Insert(27)
	heap.Insert(55)

	//         55
	//    39        41
	// 18    27  33

	if removed := heap.ExtractMax(); removed != 55 {
		t.Errorf("Expected 55, got %d", removed)
	}
	if removed := heap.ExtractMax(); removed != 41 {
		t.Errorf("Expected 41, got %d", removed)
	}
	if removed := heap.ExtractMax(); removed != 39 {
		t.Errorf("Expected 39, got %d", removed)
	}
	if removed := heap.ExtractMax(); removed != 33 {
		t.Errorf("Expected 33, got %d", removed)
	}
	if removed := heap.ExtractMax(); removed != 27 {
		t.Errorf("Expected 27, got %d", removed)
	}
	if removed := heap.ExtractMax(); removed != 18 {
		t.Errorf("Expected 18, got %d", removed)
	}
}
