package ssort

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
