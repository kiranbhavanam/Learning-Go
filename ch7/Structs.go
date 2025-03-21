// Sample program to show how to declare methods and how the Go
// compiler supports them.
package ch7

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
    name  string

    email string
}
type human struct{
	// fname string
	integer int
}
type notifier interface{
    notify()
}

func sendNotification(n notifier){
    n.notify()
}
// notify implements a method with a value receiver.
func (u *user) notify() {
    fmt.Printf("Sending User Email To %s<%s>\n",
        u.name,
        u.email)
}
func (h human) notify() {
    fmt.Printf("Sending User Email To %s<%s>\n",
        h.integer)
}
// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
    u.email = email
}
func (u user) changeEmail2(email string) {
    u.email = email
}

// main is the entry point for the application.
func Master2() {
    // Values of type user can be used to call methods
    // declared with a value receiver.
    bill := user{"Bill", "bill@email.com"}
    // sendNotification(bill)
    bill.notify()
	// h1:=&human{"kiran","kkkk"}
	// h1.notify()
	h1:=human{5}
    sendNotification(h1)
    // h1.notify()
	h1.mod()
	h1.mod2()
    
    // Pointers of type user can also be used to call methods
    // declared with a value receiver.
    lisa := &user{"Lisa", "lisa@email.com"}
    lisa.notify()
    sendNotification(lisa)
    // Values of type user can be used to call methods
    // declared with a pointer receiver.
    bill.changeEmail("bill@newdomain.com")
    bill.notify()
    // Pointers of type user can be used to call methods
    // declared with a pointer receiver.
    lisa.changeEmail2("lisa@comcast.com")
    lisa.notify()
    lisa.changeEmail("lisa@comcast.com")
    lisa.notify()
}

func (a human) mod(){
	fmt.Println(a.integer)
}
func (a *human) mod2(){
	fmt.Println(a.integer)
}
// Why this design?
//Convenience: Go reduces boilerplate by handling the conversion for you.
//Efficiency: Use value receivers when the method doesn’t modify the data. Use pointer receivers when you want to mutate the object or avoid copying large structs