package main

import "fmt"

type myNumber int

func mainadsad() {

	var myNumberData myNumber = 10
	var userDetail interface{} = "User Name Value"
	var userAge interface{} = 20
	var firstCompany interface{} = false

	fmt.Println(myNumberData)

	employeeDetails := map[string]interface{}{
		"userName":  "string value",
		"userId":    40,
		"isBoolean": true,
	}

	userDetail = 53427

	findValue(employeeDetails["userName"])
	findValue(userAge)
	findValue(userDetail)
	findValue(firstCompany)

}

func findValue(property interface{}) {
	fmt.Println(property)
}
