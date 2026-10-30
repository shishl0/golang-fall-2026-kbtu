package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func readFile(i int) []byte {
	data := make([]byte, 1<<20)
	copy(data, fmt.Sprintf("HEADER-%04d", i))
	return data
}

func keepHeadersSlice(n int) [][]byte {
	var headers [][]byte
	for i := 0; i < n; i++ {
		data := readFile(i)
		headers = append(headers, data[:11])
	}
	return headers
}

func keepHeadersCopy(n int) [][]byte {
	var headers [][]byte
	for i := 0; i < n; i++ {
		data := readFile(i)
		headers = append(headers, bytes.Clone(data[:11]))
	}
	return headers
}

func keepSubstrings(n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		big := strings.Repeat("a", 1<<20)
		result = append(result, big[:5])
	}
	return result
}

func keepClones(n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		big := strings.Repeat("a", 1<<20)
		result = append(result, strings.Clone(big[:5]))
	}
	return result
}

func liveHeapMB() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc >> 20
}

func main() {
	before := liveHeapMB()

	headers := keepHeadersSlice(100)
	fmt.Println("subslices:", liveHeapMB()-before, "MB")
	runtime.KeepAlive(headers)
	headers = nil

	copies := keepHeadersCopy(100)
	fmt.Println("copies:", liveHeapMB()-before, "MB")
	runtime.KeepAlive(copies)
	copies = nil

	subs := keepSubstrings(100)
	fmt.Println("substrings:", liveHeapMB()-before, "MB")
	runtime.KeepAlive(subs)
	subs = nil

	clones := keepClones(100)
	fmt.Println("clones:", liveHeapMB()-before, "MB")
	runtime.KeepAlive(clones)
}
