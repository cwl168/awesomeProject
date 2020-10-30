package main

import (
	"fmt"
	"runtime"
	"time"
)

func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				fmt.Println("done")
				return
			}
		}
		close(out)
	}()
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	// Set up the pipeline.
	done := make(chan struct{})
	defer close(done)
	out := gen(done, 2, 3)
	for n := range out {
		fmt.Println(n)
		time.Sleep(2 * time.Second)
		if true {
			break
		}
	}
}
