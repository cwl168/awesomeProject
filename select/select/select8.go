package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				//写入true
				o <- true
				return
			}
		}
	}()
	c <- 666 // 注释掉，引发 timeout
	<-o
}
