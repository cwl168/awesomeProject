package main

import (
	"fmt"
	"sync"
)

//并发退出 广播机制：不要向channel发送值，而是用关闭一个channel来进行广播
func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, &wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

//表示一个只接收的channel，只能接收不能发送
func doIt(workerID int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}
