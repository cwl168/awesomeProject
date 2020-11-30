package main

import (
	"fmt"
	"time"
)

func main() {
	//main goroutine 使用done信号阻塞，要是没有其他gorouine运行的话，就会发生deadlock的错误
	/*done := make(chan bool)
	go func() {
		fmt.Println("hello world")
		time.Sleep(10 * time.Second)
	}()
	<-done*/

	done := make(chan bool)
	go func() {
		for true {
			fmt.Println("hello world")
			time.Sleep(1 * time.Second)
		}
	}()
	<-done

}
