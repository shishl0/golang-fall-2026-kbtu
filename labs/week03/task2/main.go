package main

import (
	"bytes"
	"fmt"
	"strings"
)

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

func main() {
	data := []byte("\r\n\t  Hello, World!\r\n  \t")
	fmt.Printf("%q\n", processPayloadOriginal(data))
	fmt.Printf("%q\n", processPayloadRefactored(data))
}
