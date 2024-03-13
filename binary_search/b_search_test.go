package bsearch

import "testing"

func TestBSearchForValidResult(t *testing.T) {
	arr := []int32{3, 5, 7, 11, 15, 29, 30, 33, 45, 62, 82, 97}
	expectedPosition := 5

	position := BSearch(arr, 29)

	if expectedPosition != position {
		t.Error("failure to find the correctly position")
	}
}

func TestBSearchForInvalidResult(t *testing.T) {
	arr := []int32{3, 5, 7, 11, 15, 29, 30, 33, 45, 62, 82, 97}
	expectedPosition := -1

	position := BSearch(arr, 8)

	if expectedPosition != position {
		t.Error("failure to find the correctly position")
	}
}

func BenchmarkBSearch(b *testing.B) {
	arr := []int32{3, 5, 7, 11, 15, 29, 30, 33, 45, 62, 82, 97, 99, 105, 121, 144, 158, 199, 204, 250}

	for i := 0; i < b.N; i++ {
		BSearch(arr, 97)
	}
}
