package main

import (
	"fmt"
)

type User struct {
	Name, LastName string
}

func (u User) completeName() string {
	return fmt.Sprintf("Hi User %s %s", u.Name, u.LastName)
}

type Customer struct {
	Name, LastName string
}

func (c Customer) completeName() string {
	return fmt.Sprintf("Hi Customer %s %s", c.Name, c.LastName)
}

type Namer interface {
	completeName() string
}

func Greet(n Namer) string {
	return n.completeName()
}

func main() {
	u := User{"Alber", "Eisteing"}
	fmt.Println(Greet(u))
	c := Customer{"Gus", "Peiretti"}
	fmt.Println(Greet(c))
}