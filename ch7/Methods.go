package ch7

import "fmt"

type Person struct {
	Fname string
	Lname string
	Age   int
	// h Human
}

var a int
var b string

type Human struct {
	// Fname string
	// Lname string
	p Person
}
type Employee struct {
	o Person //embedded field
	// reports []Person
}
type Update interface {
	Modify()
}

func (p Person) Details() string {
	return fmt.Sprintf("%s %s,%d", p.Fname, p.Lname, p.Age)
}

func (p *Person) Modify() {
	p.Fname = "kiran"
	p.Lname = "Reddy"
	p.Age = 22
	// p.Details()
}

// func (h *Human) Modify(){
// 	// h.Fname="abc"
// 	// h.Lname="abb"
// 	// // h.Details

// }
func exec2() {
	p := Person{ //p:=&ch7.Person{
		Fname: "kkr", Lname: "bhavanam", Age: 22,
	}
	fmt.Println(p.Details()) //kkr bhavanam,22
	p.Modify()
	fmt.Println(p.Details()) //kiran Reddy,22

}
func Change(u Update) {
	// fmt.Println("u.Modify(): ")
	u.Modify()
}
func Exec() {
	emp := Employee{
		o: Person{
			"achyuth",
			"bhavanam",
			23,
		},
		// reports:[]Person{},
	}
	emp.o.Modify()
	// human:=&Human{
	// 	// Fname:"kkr",Lname:"B",
	// 	p:Person{"aaa","bbb",44,},
	// }
	// fmt.Println(emp.o)
	// Change(&human)

	// human.Modify()
	Change(&Person{})
	fmt.Println("&a: ", &a, "a: ", a)
	fmt.Println("&a: ", &b, "a: ", b)
	// fmt.Println()
	// var e1 Employee=emp //no error
	// var e2 Person=emp //error
}
