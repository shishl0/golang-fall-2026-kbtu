package main

import "testing"

var sinkN int

func BenchmarkRenderNew(b *testing.B) {
	ids := makeIDs(200)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkN = RenderNew(ids)
	}
}

func BenchmarkRenderPool(b *testing.B) {
	ids := makeIDs(200)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkN = RenderPool(ids)
	}
}
