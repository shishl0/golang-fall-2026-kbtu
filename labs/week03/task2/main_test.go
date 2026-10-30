package main

import "testing"

var testPayload = []byte("\r\n\t  Hello, World!\r\n This is a test payload with carriage returns.\r\n  \t")

func BenchmarkProcessPayloadOriginal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		processPayloadOriginal(testPayload)
	}
}

func BenchmarkProcessPayloadRefactored(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		processPayloadRefactored(testPayload)
	}
}
