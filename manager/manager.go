package manager

import emp "oop/employee"

// employee abstraction on manager
type EmployeeOnManager interface {
	SetPerformance(performance emp.Performance)
}

func Work(employees ...EmployeeOnManager) {
	for i, employee := range employees {
		employee.SetPerformance(emp.Performance(i % 3))
	}
}
