package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{5, 2, 8, 3, 1}, []int{1, 2, 3, 5, 8}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := QuickSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("QuickSort(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
