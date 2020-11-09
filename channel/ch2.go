package main

import (
	"fmt"
	"runtime"
	"time"
)

var s int

//goroutine 泄露
func counter(in chan<- int, s int) chan int {
	for i := 0; ; i++ {
		s++
		in <- i
	}
}
func main() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()

	var ch = make(chan int)
	go counter(ch, s)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("s value", s)
}
