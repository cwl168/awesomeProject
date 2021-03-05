package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//<-chan int 只能读通道
func gen(ctx context.Context) <-chan int {
	// 创建子context
	subCtx, _ := context.WithCancel(ctx)
	go sub(subCtx) // 这里使用ctx，也能给goroutine通知

	dst := make(chan int)
	n := 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("end")
				return // return，防止goroutine泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func sub(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("end too")
			return // returning not to leak the goroutine
		default:
			fmt.Println("test")
		}
	}
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	cancel()
	wg.Wait()
}
