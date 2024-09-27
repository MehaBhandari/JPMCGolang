// Multiple Producers, single Consumer Goroutines

package main

import (
	"fmt"
)

func rangeData() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 80; i++ {
			ch <- i
		}

		close(ch)
	}()

	for inputData := range ch {
		fmt.Println(inputData)
	}
}
