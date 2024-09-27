package main

import (
	"fmt"

	"github.com/MehaBhandari/GolangJPMC/structures"
)

func ImportingPackages() {
	employeeOne := structures.EmployeeDetails{}
	employeeOne.EmployeeName = "Mayank"
	employeeOne.EmployeeAddress = structures.Address{
		City:    "delhi",
		Street:  "delhi",
		Country: "India",
	}

	structures.UpdateAddress(&employeeOne.EmployeeAddress)

	employeeTwo := new(structures.EmployeeDetails)
	employeeTwo.EmployeeName = "Mayank"
	employeeTwo.EmployeeAddress = structures.Address{
		City:    "delhi",
		Street:  "delhi",
		Country: "India",
	}

	structures.UpdateData(employeeTwo)

	isEqual := (employeeOne == *employeeTwo)
	structures.UpdateAddress(&((*employeeTwo).EmployeeAddress))
	fmt.Println(isEqual)
}
