package ssort

import "testing"

func TestSelectionSort(t *testing.T) {
	unsortedVec := []int{10, 32, 5, 71, 21, 44, 55, 61, 1, 22, 89}
	sortedVec := []int{1, 5, 10, 21, 22, 32, 44, 55, 61, 71, 89}

	SelectionSort(unsortedVec)

	for i, v := range unsortedVec {
		if v != sortedVec[i] {
			t.Errorf("index: %v, expected value: %v, current value: %v", i, sortedVec[i], v)
		}
	}
}
