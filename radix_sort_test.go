package main

import (
	"reflect"
	"testing"
)

func TestValidRadixSort(t *testing.T) {
	unsorted := []int{54, 5, 9843, 6583, 784, 321, 4566, 597}
	sorted := []int{5, 54, 321, 597, 784, 4566, 6583, 9843}

	result := RadixSort(unsorted)

	if !reflect.DeepEqual(result, sorted) {
		t.Errorf("Array %d is not in the correct order", result)
	}
}
