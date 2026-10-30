package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{3, 4}
	f := r.Area
	r.Width = 10
	fmt.Println(f(), r.Area())
}
