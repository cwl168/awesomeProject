package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		//接收者的goroutine可能会或者不会有足够的时间，在发送者继续执行前处理消息。
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed  cmd.2 可能打印不出来
}
