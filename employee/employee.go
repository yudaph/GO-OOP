package employee

// performance constant
type Performance int32

const (
	Bad Performance = iota
	Good
	Great
)

type Employee struct {
	// attributes
	Name        string
	BasicSalary float64
	performance Performance
}

// employee method
func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) SetPerformance(performance Performance) {
	e.performance = performance
}

func (e *Employee) CalculateSalary() float64 {
	salary := e.BasicSalary
	salary += e.calculateBonus()
	/*
		Do more calculation
	*/
	return salary
}

func (e *Employee) calculateBonus() float64 {
	switch e.performance {
	case Great:
		return e.BasicSalary
	case Good:
		return e.BasicSalary / 2
	default:
		return 0
	}
}
