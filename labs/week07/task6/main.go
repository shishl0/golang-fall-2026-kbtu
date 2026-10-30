package main

import (
	"fmt"
	"runtime"
	"time"
)

type Record struct {
	ID   int
	Data [256]byte
}

var live []*Record
var last *Record

//go:noinline
func consume(r *Record) int {
	last = r
	return r.ID
}

func main() {
	start := time.Now()

	for i := 0; i < 80_000; i++ {
		live = append(live, &Record{ID: i})
	}

	sum := 0
	for i := 0; i < 10_000_000; i++ {
		r := &Record{ID: i}
		sum += consume(r)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("sum:", sum)
	fmt.Println("GC cycles:", m.NumGC)
	fmt.Println("heap memory from the OS:", m.HeapSys>>20, "MB")
	fmt.Println("total GC pauses:", time.Duration(m.PauseTotalNs))
	fmt.Println("time:", time.Since(start))
}
