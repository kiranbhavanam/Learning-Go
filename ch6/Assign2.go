// Write two functions. The UpdateSlice function takes in a []string and a string.
// It sets the last position in the passed-in slice to the passed-in string.
//
//	Print the slice after making the change.
//
// The GrowSlice function also takes in a []string and a string.
//
//	It appends the string onto the slice. Print the slice after making the change.
//
// Call these functions from main.
//
//	Print out the slice before each function is called and after each function is called.
package ch6

import "fmt"

func UpdateSlice(strArr []string, str string) {
	strArr[len(strArr)-1] = str
	fmt.Println(strArr)
}
func GrowSlice(strArr []string, str string) {
	strArr = append(strArr, str)
	fmt.Println(strArr)
}
