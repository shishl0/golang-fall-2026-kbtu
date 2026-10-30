package main

import "testing"

var sinkValue Point
var sinkPointer *Point
var sinkSlice []int

func BenchmarkNewPointValue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkValue = NewPointValue(1, 2)
	}
}

func BenchmarkNewPointPointer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkPointer = NewPointPointer(1, 2)
	}
}

func BenchmarkCollectNoCap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkSlice = CollectNoCap(1000)
	}
}

func BenchmarkCollectWithCap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkSlice = CollectWithCap(1000)
	}
}
