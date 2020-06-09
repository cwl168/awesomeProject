package main

import "fmt"

func main() {
	//无缓冲信道，在接收者未准备好之前，发送操作是阻塞的
	ch := make(chan string)
	go func() {
		//写
		ch <- "send"
	}()
	//读
	fmt.Println(<-ch)

}