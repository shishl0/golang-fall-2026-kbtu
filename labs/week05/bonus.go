package main

import "fmt"

func bonus() {
	var a, b any = 1, 1
	fmt.Println(a == b)

	var c, d any = 1, int64(1)
	fmt.Println(c == d)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	var e, f any = []int{1}, []int{1}
	fmt.Println(e == f)
}
