package main

import "fmt"

type DirectorInt interface {
	getSalary() int
}

type ManagerInt interface {
	getName() string
}

type IEmployee interface {
	DirectorInt
	ManagerInt
}

type Manager struct {
	id        int
	createdAt string
	name      string
	avatar    string
	salary    int
	teamSize  int
	teamName  string
}

func (m Manager) getSalary() int {
	return m.salary
}

func (m Manager) getName() string {
	return m.name
}

type voidType interface{}

type Associates struct {
	id        voidType
	createdAt any
	name      string
	avatar    string
	salary    int
}

func (m Associates) getSalary() int {
	return m.salary
}

func (m Associates) getName() string {
	return m.name
}

func main12326() {
	myAssociate := Associates{
		id:        1,
		createdAt: "2019-07-10T08:05:39.325Z",
		name:      "Kristian Stoltenberg III",
		avatar:    "https://s3.amazonaws.com/uifaces/faces/twitter/kiwiupover/128.jpg",
		salary:    10,
	}

	myManager := Manager{
		id:        1,
		createdAt: "2019-07-10T08:05:39.325Z",
		name:      "Kristian Stoltenberg III",
		avatar:    "https://s3.amazonaws.com/uifaces/faces/twitter/kiwiupover/128.jpg",
		salary:    100,
	}

	calculateData(myAssociate)
	calculateData(myManager)

	employeeArray := []IEmployee{myAssociate, myManager}

	for _, employee := range employeeArray {
		fmt.Println(employee.getName())
		fmt.Println(employee.getSalary())

		switch employee.(type) {
		case Associates:
			fmt.Println("This is an Associate")

		case Manager:
			fmt.Println("This is an Manager")
		}
	}

}

func calculateData(myObject IEmployee) {
	fmt.Println((myObject.getSalary()))
}
