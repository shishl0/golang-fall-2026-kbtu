package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Scale(k float64) Rectangle {
	return Rectangle{r.Width * k, r.Height * k}
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func task6() {
	r := Rectangle{3, 4}
	r = r.Scale(2)
	shapes := []Shape{r, &Circle{Radius: 1}}
	for _, s := range shapes {
		fmt.Println(s.Area(), s.Perimeter())
	}
}
