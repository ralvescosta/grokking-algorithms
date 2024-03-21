package ssort

func SelectionSort(arr []int) []int {
	if arr == nil || len(arr) <= 1 {
		return arr
	}

	sorted := []int{}
	copied := arr

	smallest := copied[0]
	smallestIndex := 0

	for range arr {
		for i := 0; i < len(copied); i++ {
			if copied[i] < smallest {
				smallest = copied[i]
				smallestIndex = i
			}
		}

		sorted = append(sorted, smallest)
		copied = remove(copied, smallestIndex)

		smallest = copied[0]
		smallestIndex = 0
	}

	return sorted
}

func remove(arr []int, index int) []int {
	if len(arr) <= 1 {
		return arr
	}

	return append(arr[:index], arr[index+1:]...)
}
