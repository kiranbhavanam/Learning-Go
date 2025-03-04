package ch5

import (
	"errors"
	"fmt"
	"strconv"
)

// var a int =5
// var a int=5
var a float64 = 10 //10.0

func Div(a, b int) int {
	return a / b

}

type Person struct {
	Fname string
	Lname string
	Age   int
}

func RenderDetails(p Person) {
	fmt.Println("persons name is :", p.Fname, " ", p.Lname,
		"and is of ", p.Age, " years")
	// a=8
	// a:=8

}
func AddTo(base int, values ...int) []int {
	out := make([]int, 0)
	for _, value := range values {
		out = append(out, base+value)
	}
	return out
}

func ReturnValues(a, b int) (rem int, res int, err error) {
	// rem=a/b
	// res=a%b
	if b == 0 {
		err = errors.New("denominator is 0")
		return a, b, err
	}
	return a / b, a % b, nil

}
func Fun1(a int) string {
	return strconv.Itoa(a)
}
func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func div(i, j int) int { return i / j }
func mul(i, j int) int { return i * j }
func Master() {
	// var myFunc func(int)string=Fun1
	// // myFunc=Fun1
	// res:=myFunc(55)
	// fmt.Println("res: ",res)

	opAssoc := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"/": div,
		"*": mul,
	}
	expressions := [][]string{
		{"3", "*", "6"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression:", expression)
			continue
		}
		operand := expression[1]
		opfunc, ok := opAssoc[operand]
		if !ok {
			fmt.Println("Invalid print")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println("Invalid operand:", expression[0])
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println("Invalid operand:", expression[2])
			continue
		}
		res := opfunc(p1, p2)
		fmt.Println(res)
	}

}
