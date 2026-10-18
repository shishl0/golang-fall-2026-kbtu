package main

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
