package main

import "fmt"

func main() {
	values := [][]int{}

	// These are the first two rows.
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	// Append each row to the two-dimensional slice.
	values = append(values, row1)
	values = append(values, row2)

	// Display first row.
	fmt.Println("Row 1")
	fmt.Println(values[0])

	// Display second row.
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Access an element.
	fmt.Println("First element")
	fmt.Println(values[0][0])

	values = append(values, make([]int, 0))
	values[2] = append(values[2], 7)
	values[2] = append(values[2], 8, 9)
	// Display entire slice.
	fmt.Println("Values")
	fmt.Println(values)
}
