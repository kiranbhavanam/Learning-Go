package ch6

import "fmt"

type Employees struct {
	name string
	id   int
}

func Sender3() {
	var E1 = Employees{}
	makeEmp(&E1)
	fmt.Println(E1)
	E2, _ := makeEmpWithoutPointer()
	fmt.Println(E2)
}

//Use makeEmp (with a pointer) only if modification of an existing instance is absolutely necessary.

func makeEmp(emp *Employees) error {
	emp.name = "kiran"
	emp.id = 411
	return nil
}

// Preferred approach: makeEmpWithoutPointer (return by value).
// This makes the function pure (no side effects).
// Easier to understand and use.
// Safer in concurrent scenarios
func makeEmpWithoutPointer() (Employees, error) {
	e := Employees{
		"shandhar",
		404,
	}
	return e, nil
}
