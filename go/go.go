package main

import "fmt"

//接受
func hello(pipline chan string)  {
	//读
	fmt.Println(<-pipline)
}

func main()  {
	//无缓冲信道，在接收者未准备好之前，发送操作是阻塞的  只有写端，没有读端，那么 “写端”阻塞 只有读端，没有写端，那么 “读端”阻塞
	pipline := make(chan string)
	go func() {
		//读
		fmt.Println(<-pipline)
	}()
	//写  无缓存信道 读必须在写之前
	pipline <- "hello world"

	ch := make(chan string)
	go func() {
		//写
		ch <- "send"
	}()
	//读
	fmt.Println(<-ch)
}
