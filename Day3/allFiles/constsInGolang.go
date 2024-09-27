package main

import "fmt"

func constsInGolang() {
	const age = 10
	const ageOther int = 100
	const userNameValue = "Mayank"
	const userNameValueOther string = "Mayank"

	const (
		a = "Mayank"
		b = 0
		c
	)

	const (
		x = iota
		y = iota * 4
		z = iota * 20
	)

	const (
		x1 = iota
		y1 = iota * 40
		z1 = iota * 200
		zz1
	)

	var floatValue float32 = 0
	floatValue = age
	fmt.Println(floatValue)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x + y + z)
	fmt.Println(x1 + y1 + z1 + zz1)
}
