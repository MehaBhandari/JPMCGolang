package main

import (
	"fmt"
)

func mainsdad() {
	channelDataOne, channelDataTwo := make(chan []Employees), make(chan []Employees)

	go func() {
		channelDataOne <- getEmloyeeDataFromApi("http://localhost:4000/associates")
	}()

	go func() {
		channelDataTwo <- getEmloyeeDataFromApi("http://localhost:5000/employees")
	}()

	select {
	case <-channelDataOne:
		fmt.Println("Data from 1 Received")

	case <-channelDataTwo:
		fmt.Println("Data from 2 Received")
	}
}
