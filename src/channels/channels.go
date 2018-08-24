package main

import "fmt"

func calculate(ch chan <- int, value int) {
	ch <- value * 10 // write value 10 into channel
}

func main() {

	channel := make(chan int)  // create channel

	go calculate(channel, 5)
	go calculate(channel, 10)

	v1 := <-channel // read channel
	v2 := <-channel // read channel

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Scanln() // wait, press a key
	fmt.Println("done")
}
