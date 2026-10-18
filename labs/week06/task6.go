package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Name string
	Run  func() error
}

func runTask(t Task) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("task %s: panic: %v", t.Name, r)
		}
	}()
	err = t.Run()
	if err != nil {
		return fmt.Errorf("task %s: %w", t.Name, err)
	}
	return nil
}

func SafeRun(tasks []Task) []error {
	var errs []error
	for _, t := range tasks {
		err := runTask(t)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

var ErrSentinel = errors.New("sentinel error")

func task6() {
	tasks := []Task{
		{"ok", func() error { return nil }},
		{"index", func() error {
			s := []int{1, 2, 3}
			i := 5
			fmt.Println(s[i])
			return nil
		}},
		{"nilmap", func() error {
			var m map[string]int
			m["a"] = 1
			return nil
		}},
		{"custom", func() error {
			panic("config is missing")
		}},
		{"sentinel", func() error { return ErrSentinel }},
	}

	errs := SafeRun(tasks)
	for _, err := range errs {
		fmt.Println(err)
	}
	fmt.Println(errors.Is(errors.Join(errs...), ErrSentinel))
}
