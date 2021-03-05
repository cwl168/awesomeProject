package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//判断是否有取消任务信号
func isCancelled(cancelChan chan bool) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//只要1个协程能关闭  缺点：发送信号的个数和需要关闭的协程数量必须一致，耦合性强
func cancel_1(cancelChan chan bool) {
	cancelChan <- true
}

//所有协程都能关闭  channel特性，被close之后，channel仍然可读，不但可以读取出已发送的数据，还可以不断的读取零值，但是不能往channel发送数据
func cancel_2(cancelChan chan bool) {
	close(cancelChan)
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	var wg sync.WaitGroup
	cancelChan := make(chan bool, 1)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ch chan bool) {
			for {
				if isCancelled(cancelChan) {
					fmt.Println(i, "cancelled")
				} else {
					fmt.Println(i, "done")
				}
				wg.Done()
				break
			}
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	wg.Wait()
}
