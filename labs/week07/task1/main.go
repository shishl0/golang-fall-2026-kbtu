package main

import (
	"fmt"
	"os"
)

func depth(n int) int {
	var pad [64]byte
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depth(n-1) + 1 + int(pad[(n+1)%len(pad)])
}

func depthBig(n int) int {
	var pad [1024]byte
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depthBig(n-1) + 1 + int(pad[(n+1)%len(pad)])
}

func depthNoRead(n int) int {
	var pad [1024]byte
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depthNoRead(n-1) + 1
}

func forever(n int) int {
	return forever(n+1) + 1
}

func main() {
	fmt.Println("depth:", depth(1_000_000))
	fmt.Println("depthNoRead:", depthNoRead(1_000_000))

	if len(os.Args) < 2 {
		return
	}

	if os.Args[1] == "big" {
		fmt.Println("depthBig:", depthBig(1_000_000))
	}

	if os.Args[1] == "forever" {
		defer func() {
			r := recover()
			fmt.Println("recovered:", r)
		}()
		forever(0)
	}
}
