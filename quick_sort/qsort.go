package quicksort

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	less := make([]int, 0)
	greater := []int{}

	for _, v := range array {
		if v > pivot {
			greater = append(greater, v)
		}

		if v < pivot {
			less = append(less, v)
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
