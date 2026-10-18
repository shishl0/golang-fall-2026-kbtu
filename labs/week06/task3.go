package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("student not found")

type GradeError struct {
	Student string
	Grade   int
}

func (e *GradeError) Error() string {
	return fmt.Sprintf("invalid grade %d for %s", e.Grade, e.Student)
}

type GradeBook struct {
	grades map[string]int
}

func (gb *GradeBook) Get(name string) (int, error) {
	grade, ok := gb.grades[name]
	if !ok {
		return 0, ErrNotFound
	}
	return grade, nil
}

func (gb *GradeBook) Set(name string, grade int) error {
	if grade < 0 || grade > 100 {
		return &GradeError{name, grade}
	}
	gb.grades[name] = grade
	return nil
}

func AddBonus(gb *GradeBook, name string, bonus int) error {
	grade, err := gb.Get(name)
	if err != nil {
		return fmt.Errorf("add bonus to %s: %w", name, err)
	}
	err = gb.Set(name, grade+bonus)
	if err != nil {
		return fmt.Errorf("add bonus to %s: %w", name, err)
	}
	return nil
}

func report(err error) {
	var gradeErr *GradeError
	if err == nil {
		fmt.Println("OK")
	} else if errors.Is(err, ErrNotFound) {
		fmt.Println("404:", err)
	} else if errors.As(err, &gradeErr) {
		fmt.Println("400:", err, gradeErr.Student, gradeErr.Grade)
	} else {
		fmt.Println("500:", err)
	}
}

func task3() {
	gb := &GradeBook{grades: make(map[string]int)}
	gb.Set("Ali", 90)

	report(AddBonus(gb, "Ali", 5))
	report(AddBonus(gb, "Zed", 5))
	report(AddBonus(gb, "Ali", 50))

	err := fmt.Errorf("add bonus to Ali: %v", &GradeError{"Ali", 140})
	report(err)
}
