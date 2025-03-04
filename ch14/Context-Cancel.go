package ch14

import (
	"context"
	"fmt"
	"time"
)

func ContextCancel() {

	ctx, cancel := context.WithCancel(context.Background())
	go task(ctx)
	time.Sleep(3*time.Second)
	
	cancel()
	time.Sleep(1*time.Second)
}

func task(ctx context.Context){

	for{
		select {
		case <-ctx.Done():
			fmt.Println("task terminated")
			return
		default :
			fmt.Println("task is running....")
			// time.Sleep(1*time.Second)
			
		}
	}
}