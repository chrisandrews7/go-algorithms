package main

func InsertionSort(array []int) []int {
	for i := range array[1:] {
		for index := i + 1; index > 0; index-- {
			if array[index] < array[index - 1] {
				array[index-1], array[index] = array[index], array[index-1]
			}
		}
	}

	return array
}