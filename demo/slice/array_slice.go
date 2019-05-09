package main

import "fmt"

func main() {
	var array1 [6]int
	var slice1 []int = array1[2:5] // item at index 5 not included!(5-1)
	// var slice1 []int = array1[n:m]
	// slice length is (m - n)
	// slice1 cap is 4 (array length - 2)
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of array1 is %d\n", len(array1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7] // panic: runtime error: slice bound out of range
	
	// slice as function param
	var arr = [5]int{4, 3, 5, 8, 9}
	s := arr[:]
	sum := Sum(s)
	fmt.Printf("The Sum %d\n", sum)
}

func Sum(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}
