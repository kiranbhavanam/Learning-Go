package ch4

import "fmt"

func For1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	i := 1
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	for i := 0; i < 10; {
		// fmt.Print(i)
		if i%2 == 0 {
			i++
		} else {
			i = i + 2
		}
	}
	for i < 100 {
		fmt.Print(i)
		i = i * 2
	}
	for {
		fmt.Println(i)
		if i%20 == 0 {
			fmt.Print("hello kiran!!")
			break
		}
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
	evenVals := []int{2, 4, 6, 8, 10, 12}

	// for index,value:=
	fmt.Println(evenVals)
}

func For2() {
	outerValues := []string{"kiran", "kumar"}
outer:
	for _, outerValue := range outerValues {
		for i, r := range outerValue {
			fmt.Println(i, r, string(r))
			if r == 'a' {
				continue outer
			}
		}
		fmt.Println("no skip ")

	}
}

func For3() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if v == 0 {
			continue
		}
		if v == len(evenVals)-1 {
			break
		}
		fmt.Println(i, v)
	}
}
