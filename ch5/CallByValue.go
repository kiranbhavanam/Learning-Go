package ch5

import "fmt"

type person struct {
	fname string
	age   int
	marks []int
}

func reciever(a int, b string, c map[int]int, p person) {
	a = 55
	b = "kiran"
	c[1] = 999
	p.fname = "reddy"
	p.age = 88
	p.marks[0] = 99
	p.marks[1] = 89
}
func Sender() {
	p := 43
	q := "kumar"
	r := map[int]int{1: 22, 2: 44}
	s := person{"kumar reddy", 22, []int{75, 76}}
	fmt.Println(p, q, r, s)
	reciever(p, q, r, s)
	fmt.Println(p, q, r, s)
}
