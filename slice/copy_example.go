package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2, 10)

	// Notice that copy does not enlarge destination slice
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 5, 10)
	copy(slice3, slice1)
	fmt.Println(slice3)
}
