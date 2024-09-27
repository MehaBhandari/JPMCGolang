package main

import (
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	var mutex sync.Mutex

	wg.Add(18)

	go func() {

		for inputDataFromChannel := range ch {
			// Sequential Function
			go func(inputDataFromChannel string) {
				// Asynchrous System Call

				mutex.Lock()

				file, err := os.OpenFile("fileLogger.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

				if err == nil {
					file.WriteString("the Data from the Channel is: " + inputDataFromChannel + "\n\n")
				}

				file.Close()
				wg.Done()

				mutex.Unlock()
			}(inputDataFromChannel)
		}

	}()

	go func() {
		ch <- "Data One"
		wg.Done()
	}()

	go func() {
		ch <- "Data Two"
		wg.Done()
	}()

	go func() {
		ch <- "Data Three"
		wg.Done()
	}()

	go func() {
		ch <- "Data 4"
		wg.Done()
	}()

	go func() {
		ch <- "Data 5"
		wg.Done()
	}()

	go func() {
		ch <- "Data 6"
		wg.Done()
	}()

	go func() {
		ch <- "Data 7"
		wg.Done()
	}()

	go func() {
		ch <- "Data 8"
		wg.Done()
	}()

	go func() {
		ch <- "Data 9"
		wg.Done()
	}()

	wg.Wait()

	close(ch)

}
