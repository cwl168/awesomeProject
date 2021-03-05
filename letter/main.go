package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go foo(ch)
	for d := range ch {
		fmt.Println(d)
	}
}

func foo(ch chan<- int) {
	time.Sleep(3 * time.Second)
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
