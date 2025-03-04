package ch12

import "fmt"

func worker(ch chan string) {
	fmt.Println("in worker wating for data to be passed to channel")
	ch<-"hello I wrote data into ch from worker"
	fmt.Println("worker data sent")
}

func Master(){

	ch:=make(chan string)//
	ch2:=make(chan string)
	go worker(ch)
	go worker(ch2)
	fmt.Println("in main waiting to receive:")

	message:=<-ch
	fmt.Println("data read,", message)


}

/*
The worker goroutine blocks at ch <- "Hello from worker!" until the main function reads from ch.
The program synchronizes both goroutines implicitly.
Unbuffered channels are best for strict synchronization between goroutines.
*/

/*
buffered channels:
ch:=make(chan string,3)
The worker does not block when sending 3 messages because the buffer has enough space (cap = 3).
The main function can read the messages later, avoiding immediate synchronization.
🚀 Buffered channels are useful when:

You need asynchronous communication between goroutines.
You want to avoid blocking the sender while waiting for a receiver
*/