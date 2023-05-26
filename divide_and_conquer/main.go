package main

func main() {
	println(sum([]int{1, 2, 3, 4, 5, 20}))
}

func sum(vec []int) int {
	s := 0

	if len(vec) == 0 {
		return 0
	}

	s = vec[0] + sum(vec[1:])

	return s
}
