// Multiple Producers, single Consumer Goroutines

package main

import (
	"fmt"
	"sync"
)

func mainsadasdasdqs() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func(ch chan<- int) {
		ch <- 0
		wg.Done()
	}(ch)

	go func(ch chan int) {
		ch <- 1
		wg.Done()
	}(ch)

	go func(ch chan int) {
		ch <- 2
		wg.Done()
	}(ch)

	go func(ch <-chan int) {
		for inputDataFromChannel := range ch {
			fmt.Println(inputDataFromChannel)
		}
	}(ch)

	wg.Wait()

	fmt.Println("Done")
}
