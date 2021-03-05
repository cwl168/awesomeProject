package main

import (
	"fmt"
	"runtime"
	"time"
)

//https://github.com/Shitaibin/golang_goroutine_exit

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			out = nil
			fmt.Println("producer exit")
		}()

		for i := 0; i < n; i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Millisecond * 40)
		}
	}()
	return out
}

// consumer read data from in channel, print it, and print
// all proccess count in each second
// 第一种：如果某个通道关闭后，需要退出协程，直接return即可。
func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	t := time.NewTicker(time.Millisecond * 50) //定时打印已经处理的数量
	processedCnt := 0

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
			close(finish)
		}()

		// in for-select using ok to exit goroutine
		// 协程需要从in通道读数据，还需要定时打印已经处理的数量
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Process %d\n", x)
				processedCnt++
			case <-t.C:
				fmt.Printf("Working, processedCnt = %d\n", processedCnt)
			}
		}
	}()

	return finish
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	out := producer(10)
	finish := consumer(out)

	// Wait consumer exit
	<-finish
	fmt.Println("main exit")
}
