package main

import "testing"

var ids = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var sinkS string

func TestJoinIDsSlowAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() {
		sinkS = JoinIDsSlow(ids)
	})
	t.Logf("JoinIDsSlow: %v allocs", allocs)
}

func TestJoinIDsFastAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() {
		sinkS = JoinIDsFast(ids)
	})
	if allocs > 1 {
		t.Errorf("JoinIDsFast: %v allocs, want at most 1", allocs)
	}
}

func TestJoinIDsSameResult(t *testing.T) {
	if JoinIDsSlow(ids) != JoinIDsFast(ids) {
		t.Errorf("results are different")
	}
}

func BenchmarkJoinIDsSlow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkS = JoinIDsSlow(ids)
	}
}

func BenchmarkJoinIDsFast(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkS = JoinIDsFast(ids)
	}
}
