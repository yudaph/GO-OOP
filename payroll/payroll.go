package payroll

import "fmt"

// employee abstraction on finance
type EmployeeOnPayroll interface {
	GetName() string
	CalculateSalary() float64
}

func Work(employees ...EmployeeOnPayroll) {
	for _, employee := range employees {
		fmt.Println(employee.GetName(), "salary is", employee.CalculateSalary())
	}
}
