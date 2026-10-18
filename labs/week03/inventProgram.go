package main

import "fmt"

func inventProgram() {
	inventory := map[string]int{
		"apples":  10,
		"bananas": 5,
	}
	fmt.Printf("\n%v", inventory["apples"])
	inventory["bananas"] = 12
	fmt.Printf("\n%v", inventory)
	inventory["oranges"] = 8
	fmt.Printf("\n%v", inventory)
	delete(inventory, "apples")
	fmt.Printf("\n%v", inventory)
}
