package main

import "fmt"

func Sum(x int, y int) int {
	sum := x + y
	// simulate bug
	sum = sum + 1
	//
	return sum
}

func main() {
	fmt.Println("Sum of 1 + 2 =", Sum(1, 2))
}
