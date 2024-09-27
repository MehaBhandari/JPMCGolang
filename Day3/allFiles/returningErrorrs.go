package main

import (
	"errors"
	"fmt"
	"strconv"
)

func mainwae() {
	numberValue := "3432"
	value, err := strconv.Atoi(numberValue)

	if err != nil {
		fmt.Println("Error")
		return
	}

	output, err := calculateDivision(10, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The output Number is", value)
	fmt.Println("The division is ", output)
}

func calculateDivision(a, b int) (int, error) {

	defer someCode()

	if b == 0 {
		return 0, errors.New("This is not Divisible")
	}

	return a / b, nil
}

func someCode() {
	fmt.Println("ABc")
}
