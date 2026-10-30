package main

import "testing"

func BenchmarkErrorReturn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		checkErr(-1)
	}
}

func BenchmarkPanicRecover(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		safeCheck(-1)
	}
}

func BenchmarkIncNoDefer(b *testing.B) {
	var c Counter
	for i := 0; i < b.N; i++ {
		c.IncNoDefer()
	}
}

func BenchmarkIncDefer(b *testing.B) {
	var c Counter
	for i := 0; i < b.N; i++ {
		c.IncDefer()
	}
}
