package main

import (
	"fmt"
	"sync"
)

func goRoutines() {

	var waitGroupCounter sync.WaitGroup
	var allEmployees = []Employees{}

	waitGroupCounter.Add(2)

	go func() {
		var employeesList []Employees = getEmloyeeDataFromApi("http://localhost:4000/associates")

		for _, value := range employeesList {
			allEmployees = append(allEmployees, value)
		}

		allEmployees = []Employees{}
		waitGroupCounter.Done()
	}()

	go func() {
		var employeesList []Employees = getEmloyeeDataFromApi("http://localhost:5000/employees")

		for _, value := range employeesList {
			allEmployees = append(allEmployees, value)
		}
		waitGroupCounter.Done()
	}()

	waitGroupCounter.Wait()

	for _, value := range allEmployees {
		fmt.Println(value.Name)
	}
}
