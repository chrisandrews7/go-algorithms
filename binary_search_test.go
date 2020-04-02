package main

import (
	"testing"
	"math/rand"
	"time"
)

func TestValidSearch(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	
	haystack := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	needlePosition := rand.Intn(len(haystack))

	result := BinarySearch(haystack, haystack[needlePosition])

	if result != needlePosition {
		t.Errorf("Expected position %d, not %d", needlePosition, result)
	}
}

func TestMissingSearch(t *testing.T) {
	haystack := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	result := BinarySearch(haystack, 25)

	if result != -1 {
		t.Error("Expected -1")
	}
}