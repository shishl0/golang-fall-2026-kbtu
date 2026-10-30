package main

import "fmt"

type Report struct {
	Title  string
	Values [1024]float64
}

//go:noinline
func (r Report) FirstValue() float64 {
	return r.Values[0]
}

//go:noinline
func (r *Report) FirstPointer() float64 {
	return r.Values[0]
}

type Point struct {
	X, Y float64
}

//go:noinline
func (p Point) SumValue() float64 {
	return p.X + p.Y
}

//go:noinline
func (p *Point) SumPointer() float64 {
	return p.X + p.Y
}

func main() {
	var r Report
	r.Values[0] = 1.5
	fmt.Println(r.FirstValue(), r.FirstPointer())

	p := Point{1, 2}
	fmt.Println(p.SumValue(), p.SumPointer())
}
