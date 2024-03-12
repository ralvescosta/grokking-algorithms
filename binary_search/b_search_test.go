package bsearch

import "testing"

func TestBSearch(t *testing.T) {
	arr := []int32{3, 5, 7, 11, 15, 29, 30, 33, 45, 62, 82, 97}
	var expectedPosition int32 = 5

	position := BSearch(arr, 29)

	if expectedPosition != position {
		t.Error("failure to find the correctly position")
	}
}
