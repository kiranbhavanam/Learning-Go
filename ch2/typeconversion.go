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

func Arrays(){
	var x [3]*string
	y := [3]*string{new(string), new(string), new(string)}

	fmt.Println("X: ",x)
	*y[1]="6"
	fmt.Println("Y: ",y)
	fmt.Println("Y[1]: ",*y[2])
	x=y
	fmt.Println("X: ",x)
}
func Modify(x *int){

	/*
	
	*/
	fmt.Println("x:",x)
	// a:=20
	// x=&a
	*x=10
	fmt.Println("modification call by value:",x)}

func Modify2(x int){
	fmt.Println("x:",x)
	x=10
	fmt.Println("modification call by value:",x)
}

func Ex(){
	var a int
	if(a==0){
		fmt.Println("true",a)
	}
	// a,ok:=a
	var arr [2]int
	fmt.Println(arr,arr[0],arr[1])
	if(arr[1]==0){
		fmt.Println("true")
	}
}
func Append(){
	slice := []int{10, 20, 30, 40, 50}

// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]

// Allocate a new element from capacity.
// Assign the value of 60 to the new element.
fmt.Println("OG:",slice)
fmt.Println("Copy:",newSlice)
newSlice = append(newSlice, slice...)
// newSlice=append(newSlice, 70)
// newSlice=append(newSlice, 80)

// newSlice=append(newSlice, 80)
fmt.Println("OG after mod",slice)
fmt.Println("new after mod:",newSlice)

}