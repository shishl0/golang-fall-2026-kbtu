package main

import "fmt"

type Point struct {
	X, Y int
}

//go:noinline
func NewPointValue(x, y int) Point {
	return Point{x, y}
}

//go:noinline
func NewPointPointer(x, y int) *Point {
	return &Point{x, y}
}

func CollectNoCap(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

func CollectWithCap(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	fmt.Println(NewPointValue(1, 2), *NewPointPointer(3, 4))
	fmt.Println(len(CollectNoCap(1000)), len(CollectWithCap(1000)))
}
