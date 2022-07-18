package employee

type Supervisor struct {
	Employee        // composite Employee to Supervisor
	SupervisorBonus float64
}

// Overriding CalculateSalary function from Employee Struct
func (s *Supervisor) CalculateSalary() float64 {
	salary := s.BasicSalary
	salary += s.calculateBonus()

	// add supervisor bonus to salary
	salary += s.SupervisorBonus
	/*
		Do more calculation
	*/
	return salary
}
