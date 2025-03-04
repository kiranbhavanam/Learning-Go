package ch3

import "fmt"

func MapEx() {
	var nilMap map[string]int
	zeroMap := map[string]int{}
	fmt.Println(nilMap, len(nilMap))
	fmt.Println(zeroMap, len(zeroMap))
}
func MapAcc() {
	x := map[int]string{}
	x[0] = "kiran"
	x[1] = "kumar"
	x[1] = "Reddy"
	fmt.Println(x)
	y := map[string]int{}
	y["home"] = 0
	y["mall"] = 1
	// v, ok:=y["home"]

	// fmt.Println(v)
	fmt.Println(y["shop"])
	for i, v := range x {
		fmt.Println(i, v)
	}

}
func MaoToSet() {
	intSet := map[int]bool{}
	values := []int{3, 4, 5, 6, 7, 7}
	for _, value := range values {
		intSet[value] = true // 3: true 4:true 5 : true ,,, 7 :true
	}
	fmt.Println("intSet: ", intSet, "values: ", values)
	fmt.Println(intSet[7])
}
