package ch14

import (
	"context"
	"fmt"
	"time"
)

func Metadata(){

	ctx:=context.WithValue(context.Background(),"uID",111)
	go worker(ctx)
	time.Sleep(2*time.Second)
}

func worker(ctx context.Context){
	uId:=ctx.Value("uID")
	fmt.Println("uID: ",uId)
}