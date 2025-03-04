package ch15

import (
	"testing"
	"time"
	"fmt"
	"os"
)

var testTime time.Time
func TestMain(m *testing.M) {
	testTime=time.Now()
	fmt.Println("time is : ",testTime)
	exitVal:=m.Run()
	fmt.Println("clean up code")
	os.Exit(exitVal)
	
}
func Test1(T *testing.T){
	fmt.Println("test1 uses stuff from Test main",testTime)
}
func Test2(T *testing.T){
	fmt.Println("test1 uses stuff from Test main",testTime)
}