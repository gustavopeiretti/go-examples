package main

import "fmt"

func seqOfTwo() func() int {
	i := 0
	return func() int {
		i = i + 2
		return i
	}
}

func main() {
	a := seqOfTwo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := seqOfTwo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}