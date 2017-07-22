package main

import "fmt"

func main() {
	// Create an empty slice of slices.
	animals := [][]string{}

	// Create three string slices.
	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// Append string slices to outer slice.
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	// Loop over slices in animals.
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}