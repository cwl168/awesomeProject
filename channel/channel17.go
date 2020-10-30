package main

import (
	"fmt"
	"runtime"
	"time"
)

//接收不发送
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		fmt.Println(<-ch)
	}()
}
