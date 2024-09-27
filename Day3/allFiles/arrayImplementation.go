package main

import "fmt"

func arrayImplementation() {

	numerArray := [...]int{1, 2, 3, 4, 5, 6, 7, 4, 5, 6, 7, 89, 0, 0, 5, 45, 4545, 45, 454, 54, 54545, 54, 5454}

	// Arrays are of fixed length, This operation is not allowed
	// numerArray = append(numerArray, 10)

	fmt.Println(numerArray)

	// Creating basic Array
	var arrayData [3]int
	fmt.Println(arrayData)
	arrayData[0] = 10

	//ShortHand Notation for creating array
	salaryArray := [5]float64{10.1, 23.4, 56.6, 68.5, 123}
	fmt.Println(salaryArray)

	// That the operation copies values of salaryArray in otherDetails (not by reference)
	otherDetails := salaryArray
	otherDetails[0] = 56.87
	fmt.Println(otherDetails)

	// Data is not checked for array by reference
	// 1. Check the length
	// 2 Check each value
	// That reference might be checked:
	cloneArray := salaryArray
	fmt.Println(salaryArray == cloneArray)

	// Passing Array Reference to the function
	updateArrayData(&arrayData)
	fmt.Println(arrayData[0])

}

// Pure Functionality
func updateArrayData(arrayInput *[3]int) {
	(*arrayInput)[2] = 1000
	arrayInput[1] = 453
}
