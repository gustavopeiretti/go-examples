package main

import "fmt"

func sayHello() func(string) string {
	h := ""
	return func(b string) string {
		h = h + " " +  b
		return h
	}
}

func main() {
	a := sayHello()
	b := sayHello()

	fmt.Println(a("Hello golang"))
	fmt.Println(a("how are you?"))

	fmt.Println(b("Hi!"))
	fmt.Println(b("what up?"))
}
