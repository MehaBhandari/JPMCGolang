package main

import "fmt"

func ScanAndPrint() {

	var name string = ""
	company := ""
	count := 40

	fmt.Scanln(&name)
	fmt.Scanln(&company)
	fmt.Scanln(&count)
	fmt.Println(designation)
	userSalary = 465347

	// fmt.Print("Hello World...")
	// fmt.Println("Hello All, Welcome to TechnoFunnel")
	// fmt.Println("Hello All, Welcome to TechnoFunnel")
	// fmt.Printf("Hello %s, welcome to %s. Here is the Count: %d", name, company, count)

	userDetails := fmt.Sprintf("This is %s from %s, having %d experience", name, company, count)
	fmt.Println(userDetails)
}
