package main

import (
	"fmt"
	"hello_world/ch5"
)

func floatToInt() {
	// fmt.Println("hey there kiran i am in new.go file !!")

	// var sum1 int =a+int(b)
	// var sum2 float64=float64(a)+b
	// fmt.Println(sum1,sum2)
	// fmt.Println("abc")
}
func byteToInt() {
	var a byte = 100
	var b int = 32

	var sum1 int = int(a) + b
	var sum2 byte = a + byte(b)
	fmt.Println(sum1, sum2)
}
func complexed() {
	x := complex(2.5, 3.5)
	y := complex(5.4, 7.33)
	fmt.Println(x)
	fmt.Println(y)
}
func chapter5() int {
	result := ch5.Div(5, 6)
	fmt.Println(result)
	return result
}
