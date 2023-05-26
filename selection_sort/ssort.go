package ssort

import "fmt"

func main() {
	var vec = []int{10, 32, 5, 71, 21, 44, 55, 61, 1, 22, 89}

	SelectionSort(vec)

	for _, v := range vec {
		fmt.Printf("%v - ", v)
	}

}

func SelectionSort(vec []int) {
	vecLen := len(vec)
	for i := 0; i < vecLen; i++ {
		var smallestIndex = i
		for j := i; j < vecLen; j++ {
			if vec[j] < vec[smallestIndex] {
				smallestIndex = j
			}
		}

		//swap
		store := vec[i]
		vec[i] = vec[smallestIndex]
		vec[smallestIndex] = store
	}
}
