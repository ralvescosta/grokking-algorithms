package bsearch

import (
	"log"

	"golang.org/x/exp/constraints"
)

func BSearch[T constraints.Ordered](array []T, v T) int32 {
	length := int32(len(array))
	var positionToVerify int32 = length / 2

	for {
		log.Printf("c: %v\n", positionToVerify)

		if array[positionToVerify] == v {
			return positionToVerify
		}

		if array[positionToVerify] < v {
			positionToVerify = positionToVerify + (positionToVerify / 2)
		}

		if array[positionToVerify] > v {
			positionToVerify = positionToVerify / 2
		}
	}
}
