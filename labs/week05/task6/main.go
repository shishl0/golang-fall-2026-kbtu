package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var s T
	for _, n := range nums {
		s += n
	}
	return s
}

func SumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func SumAny(nums []any) int {
	total := 0
	for _, n := range nums {
		total += n.(int)
	}
	return total
}

func ToAny(nums []int) []any {
	result := make([]any, len(nums))
	for i, n := range nums {
		result[i] = n
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(SumInts(nums), Sum(nums), SumAny(ToAny(nums)))
}
