package main

import (
	"fmt"
	"strconv"
)

func main12as() {
	numberArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 6, 7}

	iteratorFunctionReturn := getArrayIterator(0)
	iteratorFunctionReturn(numberArray)

	iteratorFunctionReturn = getArrayIterator(10)
	iteratorFunctionReturn(numberArray)
}

func getArrayIterator(userPreference int) func([]int) {
	var iteratorFunction func([]int)

	if userPreference == 0 {
		iteratorFunction = func(inputData []int) {
			fmt.Println("Even " + strconv.Itoa(len(inputData)))
		}
	} else {
		iteratorFunction = func(inputData []int) {
			fmt.Println("Odd " + strconv.Itoa(len(inputData)))
		}
	}

	return iteratorFunction
}
