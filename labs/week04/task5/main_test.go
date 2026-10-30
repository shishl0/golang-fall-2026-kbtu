package main

import "testing"

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
