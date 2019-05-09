package main

import "fmt"

func main_2() {
	var slice1 []int = make([]int, 10) // make([]int, 10, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 2 * i
	}

	// print
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice1 at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}