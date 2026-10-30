package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func apply(nums []int, f func(int) int) []int {
	result := make([]int, 0, len(nums))
	for _, n := range nums {
		result = append(result, f(n))
	}
	return result
}

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

func main() {
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

	c1 := counter()
	c2 := counter()
	fmt.Println(c1(), c1(), c1())
	fmt.Println(c2())

	double := func(n int) int {
		return n * 2
	}
	fmt.Println(apply([]int{1, 2, 3}, double))
}
