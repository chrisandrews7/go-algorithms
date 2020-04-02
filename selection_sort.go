package main

func SelectionSort(array []int) []int {
	for index, value := range array {
		minimum := index

		for nextIndex, nextValue := range array[index:] {
			if nextValue < array[minimum] {
				minimum = nextIndex + index
			}
		}
		
		array[index], array[minimum] = array[minimum], value
	}

	return array
}