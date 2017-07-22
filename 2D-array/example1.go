package main

import "fmt"

func main() {
	// Create two-dimensional array.
	letters := [2][2]string{}

	// Assign all elements in 2 by 2 array.
	letters[0][0] = "a"
	letters[0][1] = "b"
	letters[1][0] = "c"
	letters[1][1] = "d"

	// Display result.
	fmt.Println(letters)
}