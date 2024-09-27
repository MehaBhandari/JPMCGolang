package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Employees struct {
	Id             string
	Name           string
	CreatedAt      string
	Avatar         string
	SomeOtherProps int
}

func getEmloyeeDataFromApi(appUrl string) []Employees {
	var employeeArray []Employees
	response, err := http.Get(appUrl)

	if err == nil {
		byteData, _ := io.ReadAll(response.Body)
		marshalErr := json.Unmarshal(byteData, &employeeArray)
		if marshalErr != nil {
			fmt.Println(employeeArray)
		}
		return employeeArray
	}

	return []Employees{}
}

func mainasdassadasd() {

	dataChannel := make(chan []Employees)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		employeeOne := getEmloyeeDataFromApi("http://localhost:5000/employees")
		dataChannel <- employeeOne
		wg.Done()
	}()

	go func() {
		employeeTwo := getEmloyeeDataFromApi("http://localhost:4000/associates")
		dataChannel <- employeeTwo
		wg.Done()
	}()

	// Adds the data to the channel
	employeeListOne := <-dataChannel
	employeeListOne = <-dataChannel
	wg.Wait()

	// Retrieve Data from Channel
	fmt.Println(employeeListOne)

	fmt.Println("dkfhagsd")
}
