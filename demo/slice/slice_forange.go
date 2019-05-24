package main

import "fmt"

func main3() {
	slice1 := make([]int, 10) // make([]int, 10, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 2 * i
	}

	// print
	for idx, item := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", idx, item)
	}

	// 忽略index
	var item int
	for _, item = range slice1 {
		fmt.Printf("%d\n", item)
	}

	// 省略item
	for idx := range slice1 {
		fmt.Printf("%d\n", slice1[idx])
	}
}
