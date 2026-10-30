package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrBadFormat = errors.New("bad format")

type Record struct {
	Name  string
	Grade int
}

func ParseAll(input string) ([]Record, error) {
	var records []Record
	var errs []error
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			errs = append(errs, fmt.Errorf("line %d: %w", i+1, ErrBadFormat))
			continue
		}
		grade, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			errs = append(errs, fmt.Errorf("line %d: %w", i+1, err))
			continue
		}
		records = append(records, Record{strings.TrimSpace(parts[0]), grade})
	}
	return records, errors.Join(errs...)
}

func main() {
	input := `Aruzhan, 95
Dias, 8O
Timur
Bota, 99999999999999999999
Ali, 70`

	records, err := ParseAll(input)
	fmt.Println(records)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrBadFormat))
	fmt.Println(errors.Is(err, strconv.ErrSyntax))
	fmt.Println(errors.Is(err, strconv.ErrRange))

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Println(numErr.Func, numErr.Num, numErr.Err)
	}

	_, err = ParseAll("Ali, 70\nBota, 90")
	fmt.Println(err == nil)
}
