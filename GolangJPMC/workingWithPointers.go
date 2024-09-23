package main

import (
	"fmt"

	"github.com/MehaBhandari/GolangJPMC/constants"
)

func workingWithPointers() {

	constants.ShowPIValue()

	userSalary := new(int)
	*userSalary = 100
	updateAge(userSalary)
	fmt.Println(*userSalary)

	userAge := 10
	// Passing the reference of the variable - memory location
	updateAge(&userAge)
	fmt.Println(userAge) // 10
}

// userAgeRef the memory location r=that contains int
func updateAge(userAgeRef *int) {
	// Go to the memory location, using * refer the value and update.
	*userAgeRef = 50
}
