package main

import (
	"fmt"
	"time"
)

func simpleChannel() {

	dataChannel := make(chan string)

	go func() {
		time.Sleep(time.Second * 15)
		dataChannel <- "Hello World Again"
	}()

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println(<-dataChannel)
	}()

	// Adds the data to the channel
	dataChannel <- "Hello World"

	// Retrieve Data from Channel
	fmt.Println(<-dataChannel)
}
