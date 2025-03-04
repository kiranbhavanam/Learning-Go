package ch13

import (
	"fmt"
	"io"
	"strings"
)

func Reader() {

	r := strings.NewReader("hello i am kiran!!!")
	buf:=make([]byte,6)
	for{
		n,err:=r.Read(buf)
		if err==io.EOF{
			break
		}
		fmt.Println(string(buf[:n]))
	}
}