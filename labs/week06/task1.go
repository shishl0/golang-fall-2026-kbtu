package main

import "fmt"

func makeAccumulator() func(int) int {
	total := 0
	return func(n int) int {
		total += n
		return total
	}
}

func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		v, ok := cache[n]
		if ok {
			return v
		}
		v = f(n)
		cache[n] = v
		return v
	}
}

func task1() {
	acc := makeAccumulator()
	fmt.Println(acc(10), acc(5), acc(-3))

	calls := 0
	slow := func(n int) int {
		calls++
		return n * n
	}
	fast := memoize(slow)
	fmt.Println(fast(4), fast(4), fast(4), fast(5))
	fmt.Println("calls:", calls)
}
