package ch5

import (
	"fmt"
	"strconv"
)

// import "maps"

// The simple calculator program doesn’t handle one error case: division by zero.
// Change the function signature for the math operations to return both an int and an error.
// In the div function, if the divisor is 0, return errors.New("division by zero") for the error.
// In all other cases, return nil.
// Adjust the main function to check for this error.

// func add(i,j int) int{return i+j}
// func div(i,j int) int{return i/j}
// func sub(i,j int) int{return i-j}
// func mul(i,j int) int{return i*j}

func Assign1() {

	opAssoc := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"/": div,
		"*": mul,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
	}
	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Not enough args")
			continue
		}
		operand1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println("Invalid argument")
			continue
		}
		operand2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println("Invalid argument")
			continue
		}

		operatorString := exp[1]
		operator, ok := opAssoc[operatorString]
		if !ok {
			fmt.Println("not a valid operator")
		}
		res := operator(operand1, operand2)
		fmt.Println(res)

	}
}
