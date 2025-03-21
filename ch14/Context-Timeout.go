package ch14

import (
	"context"
	"fmt"
	"time"
)

//controlling timeout
//cancelling go routines
//passing meta data through go applications

//1.Controlling timeout: ✅ Setting timeouts – Stop a function if it takes too long.
//Use context.WithTimeout() to set a time limit.

func ControllTimeout(){
ctx,cancel:=context.WithTimeout(context.Background(),4*time.Second)
defer cancel()
terminate(ctx)

}

func terminate(ctx context.Context){

	select {
	case <-time.After(3*time.Second):
		fmt.Println("executed successfully")
	// case <-time.After(3*time.Second):
	// 	fmt.Println("executed successfully")
	case <-ctx.Done():
		fmt.Println("task cancelled ")
	}
}