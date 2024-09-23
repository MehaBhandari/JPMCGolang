package main

import "fmt"

func main11() {

	inputFunction := declareFunction(0)
	inputFunction()

	inputFunction = declareFunction(10)
	inputFunction()
}

// Higher Order Function - return function
func declareFunction(someInput int) func() {
	var sampleFunction func()

	if someInput == 0 {
		// Anonymous Function - Run time Execution
		sampleFunction = func() {
			fmt.Println("Hello, from Anonymous Function")
		}
	} else {
		// Anonymous Function - Run time Execution
		sampleFunction = func() {
			fmt.Println("Hello, from Specific Function")
		}
	}

	return sampleFunction
}
