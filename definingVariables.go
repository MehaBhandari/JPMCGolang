package main

import (
	"fmt"
	"reflect"
)

func definingVariables() {

	userBonus := 10
	employeeName, employeeAge := "Anshul", "10"
	employeeSalary, employeeBonus := 20, 4

	fmt.Println(userName)
	fmt.Println(userCompany)
	fmt.Println(userSalary)
	fmt.Println(userBonus)

	fmt.Println(employeeName)
	fmt.Println(employeeAge)
	fmt.Println(employeeSalary + employeeBonus)

	fmt.Printf("The variable userBonus is of type %d", reflect.TypeOf(userBonus))
}
