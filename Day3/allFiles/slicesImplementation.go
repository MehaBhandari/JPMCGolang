package main

import "fmt"

func SlicesImplementation() {
	sliceData := []int{1, 2, 3}

	otherSlice := make([]int, 1000)

	anotherSlice := sliceData
	fmt.Println(otherSlice)

	fmt.Printf("Memory address of the array: %p\n", &sliceData[0])
	fmt.Printf("Memory address of the array: %p\n", &anotherSlice[0])

	fmt.Printf("Memory address of the array: %p\n", &sliceData[0])
	fmt.Printf("Memory address of the array: %p\n", &anotherSlice[0])

	// 1. Append data to slice, existing array increase its size by double, memory location of initial array will same
	// 2. Golang create a new array with the new length and capacity

	sliceData = append(sliceData, 10)
	anotherSlice[0] = 10000

	fmt.Printf("Memory address of the array: %p\n", &sliceData[0])
	fmt.Printf("Memory address of the array: %p\n", &anotherSlice[0])

	fmt.Println(sliceData)

	fmt.Printf("Memory address of the array: %p\n", &sliceData)
	fmt.Printf("Memory address of the array: %p\n", &anotherSlice)

}
