package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}

func tourGoMapsExercise() {
	fmt.Println(WordCount("I am learning Go! I am happy"))
}
