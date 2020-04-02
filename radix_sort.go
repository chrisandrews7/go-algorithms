package main

import (
	"math"
)

func getDigitFromPosition(value int, position int) int {
	r := value % int(math.Pow(10, float64(position)))
	return r / int(math.Pow(10, float64(position-1)))
}

func sort(array []int, position int) (sorted []int) {
	digitBoxes := make([][]int, 10)

	for _, value := range array {
		digit := getDigitFromPosition(value, position)
		digitBoxes[digit] = append(digitBoxes[digit], value)
	}

	for _, box := range digitBoxes {
		for _, boxValue := range box {
			sorted = append(sorted, boxValue)
		}
	}

	return sorted
}

func RadixSort(array []int) []int {
	maxDigitsSize := array[0]
	for _, value := range array[1:] {
		size := int(len(string(value)) + 1)
		if size > maxDigitsSize {
			maxDigitsSize = size
		}
	}
	
	for i := 1; i <= maxDigitsSize; i++ {
		array = sort(array, i)
	}

	return array
}