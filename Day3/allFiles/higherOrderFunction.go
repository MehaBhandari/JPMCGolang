package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) {
	fmt.Println(a + b)
}
func multiply(a, b int) {
	fmt.Println(a * b)
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("hasgdkjasgd")
	}
	return (a / b), nil
}

// Higher Order Function
func executeOperations(a int, b int, operationFunction func(int, int) (int, error)) {
	fmt.Println("Calculating the required Results")
	returnValue, errorObj := operationFunction(a, b)

	if errorObj == nil {
		fmt.Println(returnValue)
	} else {
		fmt.Println("error")
	}
}

func higherOrderFunction() {
	//executeOperations(10, 2, sum)

	//executeOperations(10, 2, multiply)

	executeOperations(10, 0, divide)
}
