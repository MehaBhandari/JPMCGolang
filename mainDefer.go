package main

import "fmt"

// Using goto Statement
// Panic Creation and the Recovery Machanism
// Sending error from a function
// Handling Errors coming from internal Function
// Deffered statements (Expressions)

func mainDefer() {

	defer func() {
		fmt.Println("This is Error Condition")
	}()

	result := someFunctionCall()
	fmt.Printf("This is Main Function %d", result)
}

// LIFO: Last Defer is called
func someFunctionCall() int {

	defer func() {
		fmt.Println("This is Error Condition")
	}()

	userNumerator := 0
	userName := "Mayank"
	fmt.Println(userName)

	fmt.Println("This is End of the Function")

	divideData := 10 / userNumerator

	return divideData
}
