package bsearch

import (
	"golang.org/x/exp/constraints"
)

func BSearch[T constraints.Ordered](array []T, v T) int {
	var low = 0
	high := len(array) - 1

	var mid int
	var guess T

	for {
		if high < low {
			return -1
		}

		mid = (low + high) / 2
		guess = array[mid]

		if guess == v {
			return mid
		}

		if guess < v {
			low = mid + 1
			continue
		}

		if guess > v {
			high = mid - 1
			continue
		}

		low = mid + 1
	}
}
