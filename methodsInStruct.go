package main

// Function in Golang
func UpdateData(emp *Employee) {
	emp.employeeName = "TechnoFunnel"
}

type Employee struct {
	employeeName string
	employeeAge  int
}

func (emp *Employee) updateData(newAge int) {
	emp.employeeName = "Updated via Struct Methods"
	emp.employeeAge = newAge
}

func (emp *Employee) updateDataWithoutReference() {
	emp.employeeName = "Updated via Struct Methods"
}

func methodsInStruct() {
	employeeData := Employee{
		employeeName: "Mayank Gupta",
		employeeAge:  20,
	}

	UpdateData(&employeeData)

	// We need to call the method from the Object Reference
	employeeData.updateData(100)
	employeeData.updateDataWithoutReference()
}
