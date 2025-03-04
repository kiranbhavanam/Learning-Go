package ch15

import "testing"

func add(n, m int) int {
	return n + m
}

func Test_add(t *testing.T){
	r:=add(2,3)
	if r!=5{
		t.Error("Error. expected 5, returned: ",r)
	}
}