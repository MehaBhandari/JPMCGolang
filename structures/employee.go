package structures

type Address struct {
	Street  string
	City    string
	Country string
}

type EmployeeDetails struct {
	EmployeeName    string
	EmployeeAge     int
	employeeSalary  float64
	EmployeeAddress Address
}

func UpdateData(emp *EmployeeDetails) {
	emp.employeeSalary = 100000
}

func UpdateAddress(address *Address) {
	address.Country = "hdgjsb"
}
