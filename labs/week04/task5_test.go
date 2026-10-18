package main

import "testing"

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

var result float64

func BenchmarkFirstValue(b *testing.B) {
	var r Report
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = r.FirstValue()
	}
}

func BenchmarkFirstPointer(b *testing.B) {
	var r Report
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = r.FirstPointer()
	}
}

func BenchmarkPointValue(b *testing.B) {
	p := Point{1, 2}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = p.SumValue()
	}
}

func BenchmarkPointPointer(b *testing.B) {
	p := Point{1, 2}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result = p.SumPointer()
	}
}
