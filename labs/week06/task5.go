package main

import (
	"fmt"
	"time"
)

type Resource struct {
	name string
}

func Open(name string) (*Resource, error) {
	fmt.Println("open", name)
	return &Resource{name}, nil
}

func (r *Resource) Close() {
	fmt.Println("close", r.name)
}

func processAllBad(names []string) {
	for _, name := range names {
		r, err := Open(name)
		if err != nil {
			continue
		}
		defer r.Close()
	}
}

func processOne(name string) {
	r, err := Open(name)
	if err != nil {
		return
	}
	defer r.Close()
}

func processAllGood(names []string) {
	for _, name := range names {
		processOne(name)
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println(name, "took", time.Since(start))
	}
}

func deferOrder() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("defer", i)
	}
}

func deferArgs() {
	x := 1
	defer fmt.Println("x =", x)
	x = 2
}

func deferResult() (n int) {
	defer func() {
		n = n * 2
	}()
	return 21
}

func task5() {
	deferOrder()
	deferArgs()
	fmt.Println(deferResult())

	names := []string{"a", "b", "c"}
	processAllBad(names)
	processAllGood(names)

	defer timer("sort")()
	time.Sleep(10 * time.Millisecond)
}
