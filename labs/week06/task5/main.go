package main

import (
	"fmt"
	"time"
)

func lifo() {
	defer fmt.Println("1: deferred first")
	defer fmt.Println("2: deferred second")
	defer fmt.Println("3: deferred third")
	fmt.Println("main body")
}

func arguments() {
	x := 1
	defer fmt.Println("deferred x =", x)
	defer func() {
		fmt.Println("closure x =", x)
	}()
	x = 2
	fmt.Println("current x =", x)
}

func double() (result int) {
	defer func() {
		result *= 2
	}()
	return 21
}

func notDouble() int {
	result := 21
	defer func() {
		result *= 2
	}()
	return result
}

func trace(name string) func() {
	fmt.Println("enter", name)
	return func() {
		fmt.Println("exit", name)
	}
}

func work() {
	defer trace("work")()
	fmt.Println("working...")
}

func level3() {
	defer fmt.Println("defer in level3")
	panic("something went wrong deep inside")
}

func level2() {
	defer fmt.Println("defer in level2")
	level3()
	fmt.Println("after level3")
}

func level1() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered in level1:", r)
		}
	}()
	level2()
	fmt.Println("after level2")
}

func helper() {
	r := recover()
	if r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer func() {
		helper()
	}()
	panic("boom A")
}

func b() {
	defer helper()
	panic("boom B")
}

func c() {
	fmt.Println("c: recover outside defer returned", recover())
}

func tryA() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("a did not recover, caught outside:", r)
		}
	}()
	a()
}

func safeDivide(x, y int) (result int, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return x / y, nil
}

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

func main() {
	defer timer("main")()

	lifo()
	fmt.Println()
	arguments()
	fmt.Println()
	fmt.Println(double(), notDouble())
	fmt.Println()
	work()
	fmt.Println()

	level1()
	fmt.Println("main continues normally")
	fmt.Println()

	tryA()
	b()
	c()
	fmt.Println()

	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(1, 0))
	fmt.Println()

	names := []string{"a", "b", "c"}
	processAllBad(names)
	fmt.Println()
	processAllGood(names)
	fmt.Println()

	time.Sleep(10 * time.Millisecond)
}
