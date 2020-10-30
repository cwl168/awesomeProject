package main

import (
	"fmt"
	"runtime"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//发送不接收 引起的泄露
func main() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	out := gen(2, 3)
	for n := range out {
		fmt.Println(n)
		time.Sleep(5 * time.Second)
		if true {
			break
		}
	}
}
