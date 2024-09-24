package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("Error coming from Divide...")
		r := recover()

		if r != nil {
			fmt.Println("This is a small issue, you continue")
		}
	}()

	calculateDivisionOne(10, 0)
	fmt.Println("This is Further Execution")
	fmt.Println("This is Further Execution")
	fmt.Println("This is Further Execution")
	fmt.Println("This is Further Execution")
	fmt.Println("This is Further Execution")
}

func calculateDivisionOne(a, b int) (int, error) {

	_, err := os.Open("./abc.txt")
	if err != nil {
		panic("This is what I cannot handle....")
	}
	result := a / b

	return result, nil
}
