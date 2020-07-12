package main

import (
	"fmt"
)
/**
双向信道
 */
func tst(i int, page chan<- int)  {
	fmt.Printf("正在打印%d\n",i)
	page<-i
}
func main() {
	n:=4
	pipline := make(chan int)
	for i := 0; i < n; i++ {
		//4个goroutine
		go tst(i,pipline)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("第%d个答应完成\n", <-pipline)
	}


}
