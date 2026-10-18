package main

import "fmt"

func printSlice(name string, s []int) {
	fmt.Printf("  %-6s = %v  len=%d cap=%d nil=%t\n", name, s, len(s), cap(s), s == nil)
}

func carsExample() {
	cars := []string{"Ferrari", "Honda", "Ford", "BYD"}
	fmt.Println("  cars:", cars, "len", len(cars), "cap", cap(cars))

	cars = append(cars, "Toyota")
	fmt.Println("  cars:", cars, "len", len(cars), "cap", cap(cars))
}

func copyExample() {
	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)

	n := copy(dst, src)
	fmt.Println("  Copied:", dst, "Elements copied:", n)

	src[0] = 99
	fmt.Println("  Source after mod:", src)
	fmt.Println("  Dest after mod:  ", dst)
}

func task2Slices() {

	w1 := []int{10, 20, 30}
	printSlice("w1", w1)
	w2 := make([]int, 3, 5)
	printSlice("w2", w2)
	arr := [5]int{1, 2, 3, 4, 5}
	w3 := arr[1:4]
	printSlice("w3", w3)
	var w4 []int
	printSlice("w4", w4)
	w5 := []int{}
	printSlice("w5", w5)

	fmt.Print("\n---\n\n")

	base := [4]int{200, 300, 400, 500}
	view := base[1:3]

	view[0] = 999
	fmt.Println("base =", base, "view =", view)

	base[2] = 888
	fmt.Println("after base[2]=888: view =", view)

	view = append(view, 777)
	fmt.Println("base =", base, "view =", view)

	view = append(view, 555)
	fmt.Println("base =", base, "view =", view)

	fmt.Print("\n---\n\n")

	fmt.Println("Paricipation task:")
	nums := []int{1, 2, 3, 4}
	fmt.Println("Before modifySlice:", nums)
	modifySlice(nums)
	fmt.Println("After modifySlice:", nums)

	appendInside(nums)
	fmt.Println("After appendInside:", nums)

	nums = appendAndReturn(nums)
	fmt.Println("After appendAndReturn:", nums)
}

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

func appendInside(s []int) {
	s = append(s, 42)
	_ = s
}

func appendAndReturn(s []int) []int {
	return append(s, 42)
}
