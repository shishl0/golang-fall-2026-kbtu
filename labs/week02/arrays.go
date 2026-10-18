package main

import "fmt"

func task1Arrays() {
	var a [5]float64 = [5]float64{1.3, 1.5, 1.7, 1.9, 2.1}
	fmt.Println(a)

	var partial [5]int = [5]int{0: 20, 3: 40, 4: 50}
	partial2 := [5]int{10, 25, 30}
	fmt.Println(partial)
	fmt.Println(partial2)

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)

	var grid [3][2]string
	grid[0] = [2]string{"A1", "A2"}
	grid[1] = [2]string{"B1", "B2"}
	grid[2] = [2]string{"C1", "C2"}
	fmt.Println("Participation task:")
	for i, row := range grid {
		for j, val := range row {
			fmt.Printf("  grid[%d][%d] = %s ", i, j, val)
		}
		fmt.Printf("\n")
	}
}
