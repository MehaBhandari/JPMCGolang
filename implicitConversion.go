package main

import (
	"fmt"
	"reflect"
)

func implicitConversion() {
	userIntData := 10
	userDataTotalSalary := 10.5

	fmt.Println(reflect.TypeOf(userDataTotalSalary))

	// No Implicit Conversion of int to float
	userDataTotalSalary = float64(userIntData)

	// No Implicit Conversion of float to int
	someIntDta := int(userDataTotalSalary)

	// No implicit conversion while mathmatical operation
	divideOutput := userDataTotalSalary / float64(someIntDta)
	fmt.Println(divideOutput)
}
