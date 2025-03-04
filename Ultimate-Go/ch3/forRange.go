package ch3

import "fmt"

func ValueSemanticFor() {

	var values = [5]string{"kiran", "niharika", "prachi"}
	fmt.Printf("Before Modification : [%s]",values[1])
	for i, v := range values {
		values[1]="niha"
		if i==1{
			fmt.Printf("after modification : [%s]\n",v)
		}
	}
}
func PointerSemanticFor() {

	var values = [5]string{"kiran", "niharika", "prachi"}
	fmt.Printf("Before Modification : [%s]",values[1])
	for i := range values {
		values[1]="niha"
		if i==1{

			fmt.Printf("after modification : [%s]",values[1])
		}
	}
}

