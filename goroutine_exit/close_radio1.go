package main

import (
	"fmt"
	"runtime"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	N := 10
	exit := make(chan struct{})
	//done := make(chan struct{}, N)

	//start N worker goroutines
	for i := 0; i < N; i++ {
		go func(n int) {
			for {
				if isCancelled(exit) {
					fmt.Printf("worker goroutine #%d exit\n", n)
					return
				} else {
					fmt.Printf("worker goroutine #%d is working...\n", n)
				}
				//select {
				//// wait for exit signal
				//case <-exit:
				//	fmt.Printf("worker goroutine #%d exit\n", n)
				//	done <- struct{}{}
				//	return
				//case <-time.After(time.Second):
				//	fmt.Printf("worker goroutine #%d is working...\n", n)
				//}
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
	//broadcast exit signal
	close(exit)

	//wait for all worker goroutines exit
	/*for i := 0; i < N; i++ {
		<-done
	}*/

	fmt.Println("main goroutine exit")
}

/*
输出结果：
worker goroutine #0 is working...
worker goroutine #2 is working...
worker goroutine #3 is working...
worker goroutine #1 is working...
worker goroutine #4 is working...
worker goroutine #8 is working...
worker goroutine #9 is working...
worker goroutine #7 is working...
worker goroutine #6 is working...
worker goroutine #5 is working...
worker goroutine #3 is working...
worker goroutine #2 is working...
worker goroutine #0 is working...
worker goroutine #5 is working...
worker goroutine #6 is working...
worker goroutine #8 is working...
worker goroutine #9 is working...
worker goroutine #4 is working...
worker goroutine #1 is working...
worker goroutine #7 is working...
worker goroutine #3 exit
worker goroutine #8 exit
worker goroutine #7 exit
worker goroutine #1 exit
worker goroutine #9 exit
worker goroutine #5 exit
worker goroutine #6 exit
worker goroutine #0 exit
worker goroutine #2 exit
worker goroutine #4 exit
main goroutine exit

*/
