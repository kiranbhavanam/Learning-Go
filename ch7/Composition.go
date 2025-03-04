package ch7

import "fmt"

type Employer struct {
	Name   string
	Salary int
}

func (e Employer) DisplaySalary() {
	fmt.Printf("Salary of %s is %d\n", e.Name, e.Salary)
}

type Bonus struct {
	Amount int
}

type SeniorManager struct {
	Employer // Composing Employee struct
	Bonus    // Composing Bonus struct separately
}

func (sm SeniorManager) DisplayTotalCompensation() {
	total := sm.Salary + sm.Amount
	fmt.Printf("Total Compensation of %s is %d\n", sm.Name, total)
}

func main() {
	sm := SeniorManager{
		Employer: Employer{Name: "Bob", Salary: 70000},
		Bonus:    Bonus{Amount: 15000},
	}

	sm.DisplaySalary()             // From Employee struct
	sm.DisplayTotalCompensation() // Own method
}
