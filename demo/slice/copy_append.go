package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	sTo := make([]int, 10)

	n := copy(sTo, slFrom)
	fmt.Println(sTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6, 7)
	fmt.Println(sl3)
}