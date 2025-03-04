package ch13

import (
	"fmt"
	"time"
)

func Time() {

	time := time.Now()
	fmt.Println(time)
	fmt.Println("formatted time: ",time.Format("2006-01-02 15:04:05"))
}