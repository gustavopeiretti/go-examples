package main

import "fmt"

func main() {

	// slide without data
	var colors []string
	fmt.Println("Colors without data", colors)

	colors = append(colors, "red")
	fmt.Println("Colors data", colors)

	// slice
	cities := []string{"buenos aires", "sao paulo", "new york"}
	fmt.Println("Cities", cities)

	// slice using make
	m := make([]int, 5)
	fmt.Printf("Slice m - lenght=%d capacity=%d values %v", len(m), cap(m), m)
	fmt.Println()

	mc := make([]int, 0, 5)
	fmt.Printf("Slice mc - lenght=%d capacity=%d values %v", len(mc), cap(mc), mc)
	fmt.Println()

}
