package ch4

import (
	"fmt"
	"math/rand/v2"
)

func Switch1() {
	values := []string{"abdul", "adiSeshuBathula", "rohith", "kiran kumar"}
	for _, value := range values {
		switch size := len(value); size {
		case 1, 2, 3:
			fmt.Println("length is small", size)
		case 4, 5, 6, 7:
			fmt.Println("lenght is medium", size)
		case 11:
		default:
			fmt.Println("too large", size)
		}
	}

}

func Switch2() {
outer:
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			count := 1
			fmt.Println("case 0 and count is: ", count)

		case 4:
			fmt.Println("case 2,3,4")
		case 8, 9:
			fmt.Println("case using fallthrough")
		case 7:
			fmt.Println("wooooo")
			break outer
		default:
			fmt.Println("default case", i)
		}
	}
}

func SwitchFizzBuzz() {
	for i := 0; i < 50; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func Excercise() {
	x := []int{}
	var value int
	for i := 0; i < 100; i++ {
		value = rand.IntN(1000)
		x = append(x, value)
	}
	fmt.Println("X: ", x)
	for _, y := range x {
		switch {
		case y%2 == 0:
			fmt.Println("Divisible by 2", y)
		case y%3 == 0:
			fmt.Println("Divisible by 2", y)
		case y%3 == 0 && y%2 == 0:
			fmt.Println("Divisible by 6", y)
		default:
			fmt.Println("Never Mind!!", y)
		}
	}
}
func Excercise2() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
}
