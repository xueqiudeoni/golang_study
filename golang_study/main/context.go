package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	//3ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	//ctx := context.WithValue(context.Background(), "name", "swag")
	//ctx, cancel := context.WithCancel(ctx)
	message := make(chan int)
	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	defer cancel()
	time.Sleep(time.Second)
	fmt.Println("main finished")
}
func son(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值，%d\n", m)
		case <-ctx.Done():
			//fmt.Println("finished", ctx.Value("name"))
			fmt.Println("finished")
			return
		}

	}
}
