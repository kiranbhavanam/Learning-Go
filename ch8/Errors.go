package ch8

import "errors"

type Person struct {
	Fname string
	Lname string
	Age   int
	// h Human
}

//wrapping errors

// handling multiple errors.
func ValidatePerson(p Person) error {
	var err []error
	if p.Fname == "" {
		err = append(err, errors.New("Fname cant be empty"))
	}
	if p.Lname == "" {
		err = append(err, errors.New("Lname cant be empty"))

	}
	if p.Age < 0 {
		err = append(err, errors.New("age cant be empty"))

	}
	if len(err) > 0 {

		return errors.Join(err...)
	}
	return nil
}
