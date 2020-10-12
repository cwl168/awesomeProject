package main

import (
	"fmt"
	"sync"
)

var (
	counter int32          //计数器
	wg      sync.WaitGroup //信号量
)

//100万个协程来模仿100万个并发,每个协程都自增1
func main() {
	threadNum := 1000000
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go incCounter()
	}
	wg.Wait() //阻塞当前协程，直到实例中的计数器归零
	fmt.Println(counter)
}
func incCounter() {
	defer wg.Done()
	counter++
}
