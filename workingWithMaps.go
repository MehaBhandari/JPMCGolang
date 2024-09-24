package main

import "fmt"

func main4234() {

	assignDataMapData := map[bool]string{
		false: "Mayank",
		true:  "20",
	}

	fmt.Println(assignDataMapData)

	assignDataMap := map[bool]string{
		false: "Mayank",
		true:  "20",
	}

	assignDataMap[false] = "Anshul"
	dataMap := make(map[string]int)

	arr := [1]int{1}

	fmt.Println(assignDataMap[false])
	assignMap := map[[1]int]string{
		arr: "Mayank",
	}

	for key, value := range assignMap {
		for iterateKey, _ := range key {
			fmt.Println(iterateKey)
		}
		fmt.Printf("the value for key %s, is given here %d", key, value)
	}

	dataMap["mayank"] = 10
	dataMap["anshul"] = 100
	dataMap["technofunnel"] = 1000
	dataMap["meha"] = 0
	dataMap["xyz"] = 0
	fmt.Printf("The memory Location is %p ", dataMap)

	dataMap["def"] = 0
	dataMap["defsds"] = 0
	fmt.Printf("The memory Location is %p ", dataMap)
	fmt.Println()

	// Going to give me 0 and a boolean which tells whether it
	value, doExist := dataMap["xyza"]
	fmt.Println(value)
	fmt.Println(doExist)

	// Iterate through key value of the map
	for key, value := range dataMap {
		fmt.Printf("the value for key %s, is given here %d", key, value)
	}

	fmt.Printf("The Age of User Mayank is %d /n ", dataMap["mayank"])
}
