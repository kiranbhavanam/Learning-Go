package ch3

import "fmt"

func Struct1() {
	type person struct {
		name string
		age  int
		pet  string
	}
	var kiran person
	kiran = person{
		"kiran",
		22,
		"adi",
	}
	pet := struct {
		name string
		kind string
	}{
		name: "dobo",
		kind: "hush",
	}
	fmt.Println(kiran.pet)
	fmt.Println(kiran, pet)
}

// Write a program that defines a struct called Employee with three fields: firstName, lastName, and id.
// The first two fields are of type string, and the last field (id) is of type int.
// Create three instances of this struct using whatever values you’d like.
// Initialize the first one using the struct literal style without names, the second using the struct literal
//  style with names, and the third with a var declaration.
// Use dot notation to populate the fields in the third struct. Print out all three structs.

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func Employees() {

	var E3 Employee
	E1 := Employee{
		"abdul",
		"qadeer",
		400,
	}
	E2 := Employee{
		firstName: "Kiran",
		lastName:  "Bhavanam",
		id:        411,
	}
	E3.firstName = "Rohith"
	// E3.id=399
	fmt.Println("E1: ", E1)
	fmt.Println("E2: ", E2)
	fmt.Println("E3: ", E3)
}
func Ex() {

	var kiranDetails struct {
		name string
		id   int
		age  int
	}
	kiranDetails.name = "kiran"
	fmt.Println(kiranDetails)

}
