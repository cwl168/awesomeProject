package main

import (
	"fmt"
	"runtime"
	"time"
)

//模拟超时控制

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	quit := make(chan struct{})

	ch := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Printf("num:%d\n", num)
			case <-time.After(3 * time.Second):
				quit <- struct{}{}
				return
			}
		}

	}()
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	<-quit

	fmt.Println("程序结束")
}
