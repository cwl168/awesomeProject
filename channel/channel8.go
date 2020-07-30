package main

import (
	"fmt"
	"time"
)

func xrange() chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	generator := xrange()
	for i := 0; i < 1000; i++ {
		fmt.Println(<-generator)
	}
	time.Sleep(time.Second * 5)
}
