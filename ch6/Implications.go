package ch6

import "fmt"

// If the pointer passed is a nil pointer, we can not make it non nil.
// Reassignment is only possible  if there is any value assigned previously.

func Sender() {
	var x *int
	reciever(x)
	fmt.Println(x) //<nil>
}

func reciever(p *int) {
	var value = 30
	p = &value
	fmt.Println("pointer value: ", p, "actual value: ", *p) //pointer value:  0xc00000a0d8 actual value:  30
}

//If you want the value assigned to a pointer parameter to still be there when you exit the function,
// you must dereference the pointer and set the value.

func Sender2() {
	var x = 40 //oo9i
	failedUpdate(&x)
	fmt.Println("after failedUpdate function: ", x) //40
	update(&x)
	fmt.Println("after update function: ", x) //33

}

func failedUpdate(p *int) {
	var value = 33
	p = &value //iiiiuyu
}
func update(p *int) {
	var value = 33
	*p = value
} //
