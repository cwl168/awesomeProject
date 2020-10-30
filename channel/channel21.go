package main

import (
	"fmt"
	"runtime"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	c := make(chan int)
	quit := make(chan int)
	go fibonacci(c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	// close(quit)
}
