package ch12

import (
	"context"
	"errors"
	"time"
)

func Timeout[T any](worker func() T, limit time.Duration) (T,error){
	out:=make(chan T,1)
	ctx,cancel:=context.WithTimeout(context.Background(),limit)
	defer cancel()
	go func(){
		out<-worker()
	}()

	select{

	case m:=<-out:
		return m,nil
	case <-ctx.Done():
		var n T
		return n,errors.New("time out")
	}


}