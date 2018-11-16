package main

import "fmt"

func sayGoodbay() {
	fmt.Println("Goodbay")
}

func sayHello() {
	defer sayGoodbay()
	fmt.Println("Hello")
}

func main() {
	sayHello()
}