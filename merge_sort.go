package main

func merge(a []int, b []int) (merged []int) {
	iterations := len(a) + len(b)

	for index := 0; index < iterations; index++ {
		if len(b) <= 0 {
			merged = append(merged, a[0])
			a = a[1:]
		} else if len(a) <= 0 {
			merged = append(merged, b[0])
			b = b[1:]
		} else if b[0] > a[0] {
			merged = append(merged, a[0])
			a = a[1:]
		} else if b[0] < a[0] {
			merged = append(merged, b[0])
			b = b[1:]
		}
	}

	return merged
}

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])

	return merge(left, right)
}
