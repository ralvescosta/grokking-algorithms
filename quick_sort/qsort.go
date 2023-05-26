package main

// improvements: chose a random element as the pivot
func QuickSort(vec []int) []int {
	l := len(vec)
	if l < 2 {
		return vec
	}

	pivot := vec[0]
	smaller := []int{}
	grater := []int{}

	for _, v := range vec[1:] {
		if v > pivot {
			grater = append(grater, v)
		}

		if v <= pivot {
			smaller = append(smaller, v)
		}
	}

	return append(append(QuickSort(smaller), pivot), QuickSort(grater)...)
}
