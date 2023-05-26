package bsort

import (
	"testing"
)

func TestBSort(t *testing.T) {
	unsortedVec := []int{10, 20, 5, 32, 18, 27, 2}
	sortedVec := []int{2, 5, 10, 18, 20, 27, 32}

	BSort(&unsortedVec, len(unsortedVec)-1)

	for i, v := range unsortedVec {
		if v != sortedVec[i] {
			t.Errorf("index: %v, expected value: %v, current value: %v", i, sortedVec[i], v)
		}
	}
}
