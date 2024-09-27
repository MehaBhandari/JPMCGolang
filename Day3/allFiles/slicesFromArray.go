package main

import "fmt"

func SlicesFromArray() {

	// Defining a array
	arrayData := []int{}
	var arrayDataOther []int = []int{}
	arrayDataOther = append(arrayData, 10)
	arrayDataOther[0] = 345

	sliceData := arrayData[2:5] // Starts from 3, 4, 5, 6
	sliceData[0] = 3000
	fmt.Println(sliceData)
}
