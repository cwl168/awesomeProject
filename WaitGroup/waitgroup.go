package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}

		done <- true
	}()
	time.Sleep(time.Second)
	//<-done 注释 不会报死锁错误，因为主协程，子程序没有耦合关系，没有死锁是因为主协程发车走了，所以子协程也只能回家
	<-done
}