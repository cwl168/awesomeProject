package main

import (
	"fmt"
	"runtime"
	"time"
)

//goroutine 泄露
func counter2(in chan<- int, st <-chan struct{}) {
	for i := 0; ; i++ {
		//select监听信道
		select {
		case in <- i:
			fmt.Printf("push to ch %d\n", i)
		case <-st:
			fmt.Println("done")
			return
		}
	}
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	var ch = make(chan int)
	var done = make(chan struct{})
	go counter2(ch, done)
	for i := 0; i < 3; i++ {
		fmt.Printf(" %d\n", <-ch)
	}
	fmt.Println("close done")
	close(done)
	//done <- struct{}{}

}

//push to ch 0
// 0
// 1
//push to ch 1
//push to ch 2
// 2
//done
//goroutine num 1
