package main

import (
	"fmt"
	"runtime"
	"time"
)

//chanel 引起的泄露
//发送者一般都会配有相应的接收者。理想情况下，我们希望接收者总能接收完所有发送的数据，这样就不会有任何问题。但现实是，一旦接收者发生异常退出，停止继续接收上游数据，发送者就会被阻塞。
func xrange() chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)
	go func() { // 开出一个goroutine
		for i := 0; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()
	return ch
}

func main() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	generator := xrange()
	for i := 0; i < 1000; i++ { // 我们生成1000个自增的整数！
		fmt.Println(<-generator)
	}
}
