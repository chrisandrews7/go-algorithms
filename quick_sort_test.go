package main

import (
	"testing"
	"reflect"
)

func TestValidQuickSort(t *testing.T) {
	unsorted := []int{9,4,3,11,1,7,10,6,5,8,2,12}
	sorted := []int{1,2,3,4,5,6,7,8,9,10,11,12}

	result := QuickSort(unsorted)

	if !reflect.DeepEqual(result, sorted) {
		t.Errorf("Array %d is not in the correct order", result)
	}
}