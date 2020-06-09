package main

import "fmt"

var ch88 = make(chan int, 6)

func mm11() {
	for i := 0; i < 10; i++ {
		ch88 <- 8 * i
	}
	close(ch8)

}
//当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个领值
func main() {
	go mm11()
	for i := 0; i < 100; i++ {
		v := <-ch88
		fmt.Print(v, "\t")
	}

}
