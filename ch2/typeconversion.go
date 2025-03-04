package ch2

//Write a program that declares an integer variable called i with the value 20. Assign i to a floating-point variable named f. Print out i and f.
import "fmt"

func TypeCon() {
	var i = 20
	var f = i
	println(i, f)
}

//Write a program that declares a constant called value that can be assigned to
//both an integer and a floating-point variable. Assign it to an integer called i and a floating-point variable called f. Print out i and f.

func ConstCon() {
	const value = 4
	var i int = value
	var f float64 = value
	println(i, f)

}

// Write a program with three variables, one named b of type byte, one named smallI of type int32,
// and one named bigI of type uint64. Assign each variable the maximum legal value for its type; then add 1 to each variable.
// Print out their values.
func TypeConInt() {
	var b byte = 6
	var small int32 = 2147483647
	var bigI int64 = 9223372036854775807
	b = b + 1
	small = small + 1
	bigI += 1
	println(b)
	println(small)
	println(bigI)

}

func MaxofType() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}

const x int64 = 10

func Const() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1 cannot assign to constant x
	// y = "bye" cannot assgin untyped constant y

	fmt.Println(x)
	fmt.Println(y)
}
