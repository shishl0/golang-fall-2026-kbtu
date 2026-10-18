package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func stringPractice() {
	s := "世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	s = "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}

	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	s = "世界"
	runes := []rune(s)
	fmt.Print(runes[1])
	fmt.Printf("\n%c\n", runes[1])

	fmt.Println(strings.TrimRight("123oxo", "xo"))
	fmt.Println(strings.TrimSuffix("123oxo", "xo"))

	fmt.Printf("\n%v\n", newConcat([]string{"Go", "is", "an", "efficient", "and", "expressive", "language."}))
}

func newConcat(values []string) string {
	totalBytes := 0
	for _, word := range values {
		totalBytes += len(word) + 1
	}

	var builder strings.Builder

	builder.Grow(totalBytes)

	for i, word := range values {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(word)
	}
	return builder.String()
}

func processPayloadOriginal(data []byte) string {

	trimmed := string(bytes.TrimSpace(data))
	clean := strings.ReplaceAll(trimmed, "\r", "")
	return clean
}

func processPayloadRefactored(data []byte) string {
	trimmed := bytes.TrimSpace(data)
	clean := bytes.ReplaceAll(trimmed, []byte("\r"), nil)
	return string(clean)
}
