package main

import (
	"fmt"
	"runtime"
	"time"
)

//https://github.com/Shitaibin/golang_goroutine_exit

func producer(n int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		defer func() {
			close(out1)
			close(out2)
			fmt.Println("producer exit")
		}()

		for i := 0; i < n; i++ {
			fmt.Printf("send %d\n", i)
			out1 <- i
			out2 <- i
			time.Sleep(time.Millisecond * 40)
		}
	}()
	return out1, out2
}

// consumer read data from in channel, print it, and print
// all proccess count in each second
// 第二种：如果某个通道关闭了，不再处理该通道，而是继续处理其他case，退出是等待所有的可读通道关闭。我们需要使用select的一个特征：select不会在nil的通道上进行等待。这种情况，把只读通道设置为nil即可解决。
func consumer(in1 <-chan int, in2 <-chan int) <-chan struct{} {
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
		//使用,ok来退出使用for-select协程，解决是当读入数据的通道关闭时，没数据读时程序的正常结束
		for {
			select {
			case x, ok := <-in1:
				if !ok {
					in1 = nil
				}
				// Process
				fmt.Printf("Process1 %d\n", x)
				processedCnt++
			case y, ok := <-in2:
				if !ok {
					in2 = nil
				}
				// Process
				fmt.Printf("Process2 %d\n", y)
				processedCnt++
			case <-t.C:
				fmt.Printf("Working, processedCnt = %d\n", processedCnt)
			}

			// If both in channel are closed, goroutine exit
			if in1 == nil && in2 == nil {
				fmt.Println("all exit")
				return
			}
		}
	}()

	return finish
}

// 多个channel都关闭才能退出 直接 赋值 in1 = nil，in2 = nil
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	out1, out2 := producer(10)
	finish := consumer(out1, out2)

	// Wait consumer exit
	<-finish
	fmt.Println("main exit")
}
