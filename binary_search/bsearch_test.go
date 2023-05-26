package bsearch

import "testing"

func TestBinarySearch(t *testing.T) {
	sortedArray := []int{1, 5, 10, 15, 20, 32, 55, 61, 79, 110, 259, 521}
	itemToFind := 20
	expectedIndex := 4

	index := BSearch(sortedArray, itemToFind)

	if index != expectedIndex {
		t.Errorf("expected index: %v, result: %v", expectedIndex, index)
	}
}
