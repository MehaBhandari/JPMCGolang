package main

import "fmt"

func multipleReturn() {
	resultantStrValue, resultantIntValue, _ := someFunctionImplementation("User")
	resultantStrValue, _, _ = someFunctionImpl("User")
	fmt.Println(resultantStrValue)
	fmt.Println(resultantIntValue)
}

func someFunctionImplementation(inputStr string) (string, int, bool) {
	updatedData := "Hello " + inputStr
	return updatedData, 10, false
}

func someFunctionImpl(inputStr string) (string, int, bool) {
	updatedData := "Hello " + inputStr
	return updatedData, 10, false
}
