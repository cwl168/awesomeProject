package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//goroutine 泄露
func counter3(in chan<- int, ctx context.Context) {
	for i := 0; ; i++ {
		//select监听信道
		select {
		case in <- i:
			fmt.Printf("push to ch %d\n", i)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	var ch = make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go counter3(ch, ctx)
	for i := 0; i < 3; i++ {
		fmt.Printf(" %d\n", <-ch)
	}
	fmt.Println("close done")
	cancel()

}
