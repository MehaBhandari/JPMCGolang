package main

import (
	"fmt"
)

type Address struct {
	street  string
	city    string
	country string
}

type EmployeeDetails struct {
	employeeName    string
	employeeAge     int
	employeeSalary  float64
	employeeAddress Address
}

func workingWithStructs() {
	employeeOther := EmployeeDetails{}
	fmt.Println(employeeOther.employeeName)

	employeeDetailsObject := EmployeeDetails{
		employeeName:   "Mayank",
		employeeAge:    30,
		employeeSalary: 300,
	}

	employeeDetailsObject.employeeAddress = Address{
		city:    "Delhi",
		street:  "Delhi NCR",
		country: "India",
	}

	employeePointer := new(EmployeeDetails)
	employeeDataClone := employeePointer

	employeePointer.employeeName = "TechnoFunnel"
	(*employeePointer).employeeAge = 200
	(&(*employeePointer)).employeeSalary = 100000

	employeeDataClone.employeeSalary = 232434253

	changeNamePointer(employeePointer)

	fmt.Println(employeePointer)

	employeeDetailsObject.employeeAge = 100

	// It is copied by Value
	employeeClone := employeeDetailsObject
	employeeClone.employeeAge = 30076

	fmt.Print(employeeDetailsObject)
	fmt.Println(employeeClone)

	// Copied by Value
	changeName(employeeDetailsObject)
	changeNamePointer(&employeeDetailsObject)

	fmt.Print(employeeDetailsObject)
}

func changeName(emp EmployeeDetails) {
	emp.employeeName = "Random Name"
	fmt.Println(emp)
}

func changeNamePointer(emp *EmployeeDetails) {
	emp.employeeName = "Random Name Pointer"
	fmt.Println(emp)
}
