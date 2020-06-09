package main

import "fmt"

func main() {
	pipline := make(chan int, 10)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline))
	pipline<- 1
	fmt.Printf("信道中当前有 %d 个数据", len(pipline))
}