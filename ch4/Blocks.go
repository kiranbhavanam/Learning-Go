package ch4

import (
	"fmt"
	"math/rand/v2"
)

var x = 10

func Block1() {
	x := 5
	if x > 3 {
		x := 6
		fmt.Println(x)
	}
	fmt.Println("x inisde block:", x)
}
func BlockIf() {

	n := rand.IntN(10)
	fmt.Println(n)
	if n <= 2 {
		fmt.Println("Number is below 2")
	} else if n <= 5 {
		fmt.Println("number is below 6 and above 2")
	} else if n > 5 {
		fmt.Println("Number is > 5")
	}

	if k := rand.IntN(5); k%2 == 0 {
		nn := 5
		fmt.Println("Even", nn)
	} else {
		fmt.Println("Odd")
	}
}
