package ssort

func SelectionSort(arr []int) []int {
	var smallest int
	var smallestIndex int

	newArr := []int{}

	for range arr {
		smallest = arr[0]
		smallestIndex = 0

		for j := range arr {
			if arr[j] < smallest {
				smallest = arr[j]
				smallestIndex = j
			}
		}

		newArr = append(newArr, smallest)
		println(smallestIndex)
	}

	return newArr
}
