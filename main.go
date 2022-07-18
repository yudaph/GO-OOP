package main

import (
	"oop/employee"
	"oop/manager"
	"oop/payroll"
)

func main() {
	john := &employee.Employee{Name: "John", BasicSalary: 3000}
	jane := &employee.Supervisor{
		Employee:        employee.Employee{Name: "Jane", BasicSalary: 2000},
		SupervisorBonus: 500,
	}
	bane := &employee.Employee{Name: "Bane", BasicSalary: 1500}
	manager.Work(john, jane, bane)
	payroll.Work(john, jane, bane)
	/*
	   John salary is 3000
	   Jane salary is 3500
	   Bane salary is 3000
	*/
}
