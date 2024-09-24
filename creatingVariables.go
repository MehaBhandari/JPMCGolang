package main

import (
	"fmt"
	"reflect"
)

var userName string = "Mayank"
var userCompany = "TechnoFunnel"
var userSalary = 100

var (
	empName, empDesignation        string
	empSalary, empBonus, userBonus int
)

var (
	name, designation  = "Mayank", "10"
	salary, bonus      = 10, 1
	isFresher, inIndia = false, true
)

func creatingVariables() {

	userBonus := 10
	employeeName, employeeAge := "Anshul", 19
	employeeSalary, employeeBonus := 20, 4

	fmt.Println(userName)
	fmt.Println(userCompany)
	fmt.Println(userSalary)
	fmt.Println(userBonus)

	fmt.Println(employeeName)
	fmt.Println(employeeAge)
	fmt.Println(employeeSalary + employeeBonus)

	fmt.Printf("The variable userBonus is of type %d", reflect.TypeOf(userBonus))

	someFunction()
}

func someFunction() {
	fmt.Println(userBonus)
	fmt.Println(salary)
	fmt.Println(bonus)
}
