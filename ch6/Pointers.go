package ch6

import "fmt"

func Pointer1() {
	var x int
	var pointerX = &x
	fmt.Println("Pointer of X", pointerX)
	fmt.Println("value of pointer,*pointerX", *pointerX)
	fmt.Println("Details of kiran using a new string varaible", kiran)
	fmt.Println("Details of abdul using a function", abdul)
}

type Person struct {
	fname string
	lname *string
	age   *int
}

var lnameStr = "kumar"
var kiran = Person{
	"kiran",
	// &lnameStr,//"kumar" here gives error expecting *string, &"kumar" untyped string constant.
	makePonter("kumar"),
	// &"k",
	makePonter(22),
}

func makePonter[T any](t T) *T {
	return &t
}

var abdul = Person{
	"abdul",
	makePonter("Qadeer"), //func copies the const as a parameter,which is a variable.
	//a variable has address and the address is returned.
	makePonter(22),
}
