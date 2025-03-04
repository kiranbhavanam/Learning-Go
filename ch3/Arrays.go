package ch3

import "fmt"

// import "slices"
var x = []int{4, 5, 6, 6}

// var y =[]int{}
func Printer() {
	fmt.Println(x, len(x), cap(x))
	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))
	// var x byte=8
	// var y =x
	// x=append(x,3)
	// x:=make([] int,2,20)
	// x=append(x,2,4,5,6,6)
	x := []int{3, 4} //error because previously x is int, only int types.
	fmt.Println(x, len(x), cap(x))

	// fmt.Println(len(y))
	// fmt.Println(slices.Equal(x,y))
	// y=append(y,x...)
	// fmt.Println(slices.Equal(x,y))
	// fmt.Println(len(x))

	// var x=[]int{3,4,5}
	// fmt.Println(x,len(x))
	// clear(x)
	// fmt.Println(x,len(x))
	// 	[3 4 5] 3
	// [0 0 0] 3
}

// var x int

func SliceOP() {
	x := []int{} //false - Zero length slice!=nil
	fmt.Println("x==nil", x == nil)
	var y []int //true- creates a nil
	fmt.Println("y==nil", y == nil)
	var z = []int{}
	fmt.Println("z==nil", z == nil)
	// x:=[]int{4,4,5,6}
	// y:=x
	// fmt.Println(x==y) can't compare two slices.
	//slicing
	x = []int{1, 999, 3, 4, 6, 7} // cannort readeclare,has to be atleast 1 new var before :=
	y = x[2:]
	z = x[:]
	//a=x[1:4] error -> no such variable a
	a := x[1:4] //[1,999,33,4]
	//b:=x[:9]//panic: runtime error: slice bounds out of range [:9] with capacity 7
	b := x[:6]
	x[1] = 999
	y[1] = 989
	z[3] = 000
	fmt.Println("X|: ", x) //changes are reflected in all the slices.
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

}

func SliceCP() {
	x := []int{4, 5, 6, 7}
	y := make([]int, 3)
	count := copy(y, x[:3]) //[4,5,6]
	// count:=copy(x[1:],x[:3]) //x=[4 4 5 6]
	fmt.Println(y, x, count)
}

func Strings() {
	x := "Kiran"
	y := x[4]
	fmt.Println(y) //retursn the ascii of n
	z := x[3:]
	fmt.Println(z)

}
func SliceToArray() {
	x := []int{1, 2, 3, 4}
	y := [3]int(x)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	x[0] = 999
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
func SliceAppend() {
	x := []int{2, 3, 4, 5}
	y := x[:3]
	z := x[1:]
	y = append(y, 999)
	y = append(y, 000)
	z = append(z, 777)
	z = append(z, 888)
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)
	fmt.Println("X: ", x)

}
func ArrayToSlice() {
	xArray := [4]int{3, 4, 5, 6}
	ySlice := xArray[:]
	fmt.Println("xArray:", xArray)
	fmt.Println("ySlice:", ySlice)
	xArray[0] = 999
	fmt.Println("xArray:", xArray)
	fmt.Println("ySlice after modification of x :", ySlice)
}
func StringRune() {
	x := "Hi 👩 and👨"
	runes := []rune(x)
	fmt.Println(len(x), len(runes))
	fmt.Println(runes)
	fmt.Println(x[:4])

	char := runes[3]
	fmt.Println(string(char))

}
