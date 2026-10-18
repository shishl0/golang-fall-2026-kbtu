package main

import (
	"errors"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
}

func describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int, int64:
		return fmt.Sprintf("integer %v", x)
	case string:
		return fmt.Sprintf("string of %d bytes", len(x))
	case []int:
		return fmt.Sprintf("slice with %d ints", len(x))
	case bool:
		return fmt.Sprintf("bool %v", x)
	case map[string]int:
		return fmt.Sprintf("map with %d keys", len(x))
	case error:
		return "error: " + x.Error()
	case fmt.Stringer:
		return "Stringer: " + x.String()
	default:
		return fmt.Sprintf("unknown type %T", x)
	}
}

func describeNoSwitch(v any) string {
	if v == nil {
		return "nil"
	}
	if x, ok := v.(int); ok {
		return fmt.Sprintf("integer %v", x)
	}
	if x, ok := v.(int64); ok {
		return fmt.Sprintf("integer %v", x)
	}
	if x, ok := v.(string); ok {
		return fmt.Sprintf("string of %d bytes", len(x))
	}
	if x, ok := v.([]int); ok {
		return fmt.Sprintf("slice with %d ints", len(x))
	}
	if x, ok := v.(bool); ok {
		return fmt.Sprintf("bool %v", x)
	}
	if x, ok := v.(map[string]int); ok {
		return fmt.Sprintf("map with %d keys", len(x))
	}
	if x, ok := v.(error); ok {
		return "error: " + x.Error()
	}
	if x, ok := v.(fmt.Stringer); ok {
		return "Stringer: " + x.String()
	}
	return fmt.Sprintf("unknown type %T", v)
}

type Timeout struct {
	Sec int
}

func (t Timeout) Error() string {
	return fmt.Sprintf("timeout after %ds", t.Sec)
}

func (t Timeout) String() string {
	return "Timeout{" + fmt.Sprint(t.Sec) + "}"
}

func kind(v any) string {
	switch v.(type) {
	case fmt.Stringer:
		return "Stringer"
	case error:
		return "error"
	case Timeout:
		return "Timeout"
	}
	return "other"
}

func task2() {
	values := []any{
		nil,
		7,
		int64(7),
		"привет",
		[]int{1, 2, 3},
		errors.New("disk full"),
		Celsius(21.5),
		3.14,
		struct{}{},
		true,
		map[string]int{"a": 1},
	}
	for _, v := range values {
		fmt.Println(v, "->", describe(v))
		fmt.Println(v, "->", describeNoSwitch(v), "(no switch)")
	}

	fmt.Println("\n--- Timeout question ---")
	fmt.Println(kind(Timeout{5}))
	var e error = Timeout{3}
	fmt.Println(e)
}
