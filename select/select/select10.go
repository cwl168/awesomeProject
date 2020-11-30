package main

import "fmt"

func main() {
	//主协程没受到死锁的影响所以不会报死锁错误
	/*ch := make(chan string)
	go func() {
		ch <- "send"
	}()
	time.Sleep(time.Second * 3)*/

	/*ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch2 <- "ch2 value"
		ch1 <- "ch1 value"
	}()

	go func() {
		<-ch1
		<-ch2
	}()

	time.Sleep(time.Second * 2)*/

	pipline := make(chan string)
	pipline <- "hello world"
	fmt.Println(<-pipline)

}
