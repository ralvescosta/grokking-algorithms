package ssort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 5, 2, 4, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := SelectionSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SelectionSort(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
