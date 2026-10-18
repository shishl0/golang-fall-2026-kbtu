package main

import (
	"cmp"
	"fmt"
)

type Score int

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

func Average[T Number](nums []T) (float64, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	total := float64(Sum(nums))
	count := float64(len(nums))
	return total / count, true
}

func MaxOf[T cmp.Ordered](items []T) (T, bool) {
	var zero T
	if len(items) == 0 {
		return zero, false
	}
	biggest := items[0]
	for _, v := range items {
		if v > biggest {
			biggest = v
		}
	}
	return biggest, true
}

func IndexOf[T comparable](items []T, target T) int {
	for i, v := range items {
		if v == target {
			return i
		}
	}
	return -1
}

func Filter[T any](items []T, keep func(T) bool) []T {
	var result []T
	for _, v := range items {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func isPassed(s Score) bool {
	return s >= 75
}

func task5() {
	scores := []Score{90, 60, 75, 82, 40}

	fmt.Println("Sum ints:", Sum([]int{1, 2, 3}))
	fmt.Println("Sum floats:", Sum([]float64{1.5, 2.5}))
	fmt.Println("Sum scores:", Sum(scores))

	fmt.Println(Average(scores))
	fmt.Println(Average([]Score{}))

	fmt.Println(MaxOf([]int{4, 9, 2}))
	fmt.Println(MaxOf([]string{"apple", "banana", "cherry"}))
	fmt.Println(MaxOf([]int{}))
	fmt.Println(MaxOf(scores))

	fmt.Println("IndexOf 75:", IndexOf(scores, 75))
	fmt.Println("IndexOf 100:", IndexOf(scores, 100))

	passed := Filter(scores, isPassed)
	fmt.Println("passed:", passed)
}
