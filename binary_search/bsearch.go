package bsearch

func BSearch(array []int, n int) int {
	begin := 0
	end := len(array) - 1

	for {
		if begin > end {
			break
		}

		mid := (begin + end) / 2
		guess := array[mid]

		if guess == n {
			return mid
		}

		if guess > n {
			end = mid + 1
			continue
		}

		begin = mid - 1
	}

	return -1
}
