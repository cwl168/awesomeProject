package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
		go incCounter(i)
	}
	wg.Wait()
	fmt.Println(counter)
}
func incCounter(index int) {
	defer wg.Done()
	atomic.AddInt32(&counter, 1)
}
