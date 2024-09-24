package main

import "fmt"

func main1() {
	intData, _ := myVariadicFunction(1, 2, 3, 4, 5, 6, 48, 24)
	fmt.Print(intData)
}

func myVariadicFunction(inputParams ...int) (int, string) {
	sum := 0
	for _, value := range inputParams {
		sum += value
	}

	return sum, "Mayank"
}
