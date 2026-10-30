package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "invalid field: " + e.Field
}

func validateBuggy(age int) error {
	var err *ValidationError
	if age < 0 {
		err = &ValidationError{Field: "age"}
	}
	return err
}

func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}
	return nil
}

func main() {
	fmt.Println("--- buggy ---")
	err := validateBuggy(20)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("err == nil:", err == nil)
	fmt.Printf("type of err: %T\n", err)

	fmt.Println("--- fixed ---")
	err = validate(20)
	fmt.Println("validate(20) == nil:", err == nil)

	err = validate(-1)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("errors.As got field:", ve.Field)
	}
}
