package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10
	loopStartedAt := time.Now()
	loopResult := FibonacciLoop(n)
	loopEndedAt := time.Now()

	memorizationStartedAt := time.Now()
	memory := map[int]int{}
	memorizationResult := FibonacciRecursivelyWithMemorization(n, &memory)
	memorizationEndedAt := time.Now()

	fmt.Printf("ooo: %v\n\n", Fib3(10))

	println(fmt.Sprintf("FibonacciLoop of %v is %v, took = %vns", n, loopResult, loopEndedAt.Nanosecond()-loopStartedAt.Nanosecond()))
	println(fmt.Sprintf("FibonacciRecursivelyWithMemorization of %v is %v, took = %vns", n, memorizationResult, memorizationEndedAt.Nanosecond()-memorizationStartedAt.Nanosecond()))
}

func FibonacciRecursively(f int) int {
	if f < 2 {
		return f
	}

	return FibonacciRecursively(f-1) + FibonacciRecursively(f-2)
}

func FibonacciRecursivelyWithMemorization(f int, memory *map[int]int) int {
	if _, ok := (*memory)[f]; ok {
		return (*memory)[f]
	}

	if f <= 2 {
		return 1
	}

	(*memory)[f] = FibonacciRecursivelyWithMemorization(f-2, memory) + FibonacciRecursivelyWithMemorization(f-1, memory)
	return (*memory)[f]
}

func FibonacciLoop(f int) int {
	current := 0
	last := 0

	for i := 1; i <= f; i++ {
		if i == 1 {
			current = 1
			last = 0
		} else {
			current += last
			last = current - last
		}
	}

	return current
}

func Fib(n int) int {
	if n <= 2 {
		return 1
	}

	return Fib(n-2) + Fib(n-1)
}

func Fib2(n int) int {
	newResult := 0
	lastResult := 0

	for i := 0; i <= n; i++ {
		if i == 1 {
			newResult = 1
			lastResult = 0
			continue
		}
		newResult += lastResult
		lastResult = newResult - lastResult
	}

	return newResult
}

func Fib3(n int) int {
	lastResult := 0
	accumulator := 0

	for i := 0; i <= n; i++ {
		if i == 1 {
			accumulator = 1
			// lastResult = 0
			continue
		}

		accumulator += lastResult
		lastResult = accumulator - lastResult
	}

	return accumulator
}
