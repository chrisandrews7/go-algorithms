package main

const PIVOT_POINT = 0

func pivot(array []int) (pivotIndex int) {
	for i, value := range array[PIVOT_POINT+1:] {
		index := i + 1
		if value < array[PIVOT_POINT] {
			pivotIndex++
			array[pivotIndex], array[index] = value, array[pivotIndex]
		}
	}
	array[pivotIndex], array[PIVOT_POINT] = array[PIVOT_POINT], array[pivotIndex]

	return pivotIndex
}

func QuickSort(array []int) []int {
	if (len(array) <= 1) {
		return array
	}

	pivotIndex := pivot(array)
	QuickSort(array[:pivotIndex])
	QuickSort(array[pivotIndex+1:])
	
	return array
}