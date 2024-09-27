package main

type DirectorInterface interface {
	showData() int
}

type ManagerInterface interface {
	showTeam() int
}

type IEmployeeData interface {
	DirectorInterface
	ManagerInterface
}
