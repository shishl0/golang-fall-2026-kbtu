package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func double(n *int) {
	*n = *n * 2
}

func movePoint(p Point) {
	p.X = p.X + 100
}

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println(math.Pi)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)

	x := 42
	var y int64 = 7
	fmt.Println(x + int(y))

	a, c := swap("hello", "world")
	fmt.Println(a, c)
	fmt.Println(split(17))

	n := 10
	double(&n)
	fmt.Println(n)

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	if v := math.Sqrt(16); v > 3 {
		fmt.Println(v)
	}

	day := 3
	switch day {
	case 1, 7:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	for i := 0; i < 3; i++ {
		defer fmt.Println("deferred", i)
	}

	p := Point{1, 2}
	movePoint(p)
	fmt.Println(p)

	q := p
	q.X = 50
	fmt.Println(p, q)

	arr := [3]int{1, 2, 3}
	arr2 := arr
	arr2[0] = 99
	fmt.Println(arr, arr2)
}
