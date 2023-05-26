package main

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	sorted := QuickSort([]int{10, 4, 22, 1, 29})

	expected := []int{1, 4, 10, 22, 29}

	for i, v := range sorted {
		if v != expected[i] {
			t.Errorf("index: %v, expected value: %v, current value: %v", i, expected[i], v)
		}
	}
}
