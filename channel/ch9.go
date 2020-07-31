package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	ch:=make(chan int)
	go pump(ch)
}
//因为没有接受，发送那里阻塞的
func pump(ch chan <- int)  {
	ch<-1
	fmt.Println("pump")
}
