package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
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
	fmt.Println(strings.TrimLeft("oxo123", "xo"))
	fmt.Println(strings.TrimPrefix("oxo123", "xo"))

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
