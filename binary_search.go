package main

func BinarySearch(haystack []int, needle int) int {
	var left int
	var middle int
	var right int = len(haystack) - 1

	for haystack[middle] != needle && left <= right {
		middle = (right + left) / 2

		if haystack[middle] < needle {
			left = middle + 1
		} else if haystack[middle] > needle {
			right = middle
		} 
	}

	if haystack[middle] != needle {
		return -1
	}

	return middle
}