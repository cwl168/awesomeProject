package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//goroutine 泄露
func counter4(in chan<- int, ctx context.Context) {
	for i := 0; ; i++ {
		//select监听信道
		select {
		case in <- i:
			fmt.Printf("push to ch %d\n", i)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			close(in)
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
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	go counter4(ch, ctx)
	for v := range ch {
		fmt.Println(v)
	}

}
