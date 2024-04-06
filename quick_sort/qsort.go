package quicksort

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivotIndex := r.Intn(len(array))
	pivot := array[pivotIndex]

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
