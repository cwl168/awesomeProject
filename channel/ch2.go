package main

import (
	"fmt"
	"runtime"
	"time"
)

//goroutine 泄露
func counter(in chan<- int) chan int {
	for i := 0; ; i++ {
		in <- i
	}
}

func main() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	var ch = make(chan int)
	go counter(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
