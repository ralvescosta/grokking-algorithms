package bsearch

// func main() {
// 	println(fmt.Sprintf("position: %v", BSearch([]int{1, 5, 10, 15, 20, 32, 55, 61, 79, 110, 259, 521}, 20)))
// 	v1 := [][]int{}

// 	v1[1][1] = 0

// 	println(v1)
// }

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
