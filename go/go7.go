package main

import (
	"time"
)
//不会报错 ，从执行结果来看，并没有子协程因为一直阻塞就造成报死锁错误，主协程，子协程 两者没有耦合的关系。
func main() {
	ch := make(chan string)
	go func() {
		//阻塞
		ch <- "send"
	}()
	time.Sleep(time.Second * 3)
}
