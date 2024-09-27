package main

import (
	"fmt"
	"time"
)

func mainasdasd() {

	dataChannel := make(chan int, 2)
	dataChannel <- 10
	dataChannel <- 20

	go func() {
		time.Sleep(time.Second * 5)
		dataChannel <- 50
	}()

	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)

	fmt.Println("Hello All")
}
