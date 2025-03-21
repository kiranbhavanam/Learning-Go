package ch12

import (
	"fmt"

)
func goroutine(ch1 chan int,ch2 chan string){
	fmt.Println("inside go routine: ")
	value:=<-ch1
	fmt.Println("read value from ch1",value)
	ch2<-"kiran"

}
func Deadlock() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go goroutine(ch1,ch2)
	fmt.Println("Main method started: ")
	select{
	case ch1<-22:
		fmt.Println("writing into ch1: ")
	case message:=<-ch2:
		fmt.Println("reading done form ch2, ",message)
		
	}
	// ch1<-22
	// fmt.Println("writing into ch1 is done, reading from ch2:")
	// message:=<-ch2
	// fmt.Println("reading done form ch2, ",message) //generates a deadlock
	
}