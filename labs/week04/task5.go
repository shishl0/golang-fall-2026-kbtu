package main

import "fmt"

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

func task5() {
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
