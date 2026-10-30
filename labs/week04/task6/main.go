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

func shapes() {
	r := Rectangle{3, 4}
	r = r.Scale(2)
	shapes := []Shape{r, &Circle{Radius: 1}}
	for _, s := range shapes {
		fmt.Println(s.Area(), s.Perimeter())
	}
}

type Speaker interface {
	Speak() string
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": meow"
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return d.Name + ": woof"
}

type Base struct{}

func (Base) Hello() string {
	return "hello"
}

func (*Base) Speak() string {
	return "speak"
}

type ByValue struct {
	Base
}

type ByPointer struct {
	*Base
}

type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return b.Title + " by " + b.Author
}

func methodSets() {
	var s Speaker
	s = Cat{"Murka"}
	fmt.Println(s.Speak())
	s = &Cat{"Barsik"}
	fmt.Println(s.Speak())
	s = &Dog{"Rex"}
	fmt.Println(s.Speak())

	s = &ByValue{}
	fmt.Println(s.Speak())
	s = ByPointer{Base: &Base{}}
	fmt.Println(s.Speak())
	s = &ByPointer{Base: &Base{}}
	fmt.Println(s.Speak())

	b := Book{"The Go Programming Language", "Donovan & Kernighan"}
	fmt.Println(b)
	fmt.Println(&b)
}

func main() {
	shapes()
	fmt.Println()
	methodSets()
}
