package main

import (
	"fmt"
	"time"
)

//向已关闭的Channel发送会引起Panic
func main() {
	ch := make(chan int)
	for i := 0; i < 2; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get the first result
	fmt.Println(<-ch)
	close(ch) //not ok (you still have other senders)
	//do other work
	time.Sleep(2 * time.Second)
}
