package main

import "fmt"

func calc(ch chan <- int, value int) {
	ch <- value * 10 // write value 10 into channel
}

func main() {

	channel := make(chan int)  // create channel

	go calc(channel, 10)
	go calc(channel, 20)
	go calc(channel, 30)

	v1 := <-channel // read channel
	v2 := <-channel // read channel
	v3 := <-channel // read channel

	fmt.Println("Values", v1, v2, v3)

	//fmt.Scanln() // wait, press a key
	//fmt.Println("done")
}
