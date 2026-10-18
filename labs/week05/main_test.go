package main

import "testing"

func makeNums() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	return nums
}

func BenchmarkSumInts(b *testing.B) {
	nums := makeNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumInts(nums)
	}
}

func BenchmarkSumGeneric(b *testing.B) {
	nums := makeNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum(nums)
	}
}

func BenchmarkSumAny(b *testing.B) {
	anyNums := ToAny(makeNums())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumAny(anyNums)
	}
}

func BenchmarkSumAnyWithBoxing(b *testing.B) {
	nums := makeNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumAny(ToAny(nums))
	}
}
