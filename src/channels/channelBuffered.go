package main

import "fmt"

func addOneHundred(message chan <- int, value int) {
	message <- value + 100
}

func main() {
	channel := make(chan int, 2)  // create channel

	go addOneHundred(channel, 100) // send value 100 to channel
	go addOneHundred(channel, 200) // send value 200 to channel
	go addOneHundred(channel, 300) // send value 200 to channel

	v1 := <-channel // read channel
	//v2 := <-channel // read channel - not mandatory
	//v3 := <-channel // read channel - not mandatory

	fmt.Println("Value", v1)
	//fmt.Println("Value", v2)
	//fmt.Println("Value", v3)
}
