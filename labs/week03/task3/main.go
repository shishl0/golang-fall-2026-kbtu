package main

import "fmt"

func main() {

	m := map[string]int{"one": 1, "two": 2}
	fmt.Printf("\n%v", m)
	fmt.Printf("\n%v", m["one"])
	m["two"] = 20
	fmt.Printf("\n%v", m)
	delete(m, "one")
	fmt.Printf("\n%v", m)

	val, ok := m["one"]
	if ok {
		fmt.Printf("\nKey exists with value: %d", val)
	} else {
		fmt.Println("\nKey not found")
	}

	newmap := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	fmt.Printf("\n%v", newmap)

	bigMap := make(map[string]int, 1_000_000)
	fmt.Printf("\n%v", len(bigMap))

	scores := map[string]int{
		"Alice": 90,
		"Bob":   0,
	}

	val, ok = scores["Alice"]
	if ok {
		fmt.Printf("\nAlice is in the map with score: %d", val)
	} else {
		fmt.Println("\nAlice not found")
	}

	val, ok = scores["Bob"]
	if ok {
		fmt.Printf("\nBob is in the map with score: %d", val)
	} else {
		fmt.Println("\nBob not found")
	}

	val, ok = scores["Charlie"]
	if ok {
		fmt.Printf("\nCharlie is in the map with score: %d", val)
	} else {
		fmt.Println("\nCharlie not found")
	}

}
