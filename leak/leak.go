package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func Test3() {
	c := make(chan struct{})
	count := 0
	//time.After()表示time.Duration长的时候后返回一条time.Time类型的通道消息。
	timeout := time.After(time.Second * 6)
	done1 := make(chan struct{})

	go func() {
		for {
			select {
			case c <- struct{}{}: //往通道写值
				time.Sleep(time.Millisecond * 100)
			case <-done1:
				log.Println("退出 子goroutine...")
				return
			}
		}

	}()
	for {
		select {
		case <-c:
			count++
			fmt.Printf("count:%d\n", count)
		case <-timeout:
			fmt.Println("超时")
			return
		default:
			time.Sleep(time.Millisecond * 50)
			if count > 3 {
				done1 <- struct{}{}
				log.Println("退出 主函数...")
				return
			}
			fmt.Println("default")
		}
	}
}

func main() {
	go func() {
		for {
			log.Println("runtime.NumGoroutine : ", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 500)
		}
	}()
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 10; i++ {
		go Test3()
	}
	time.Sleep(time.Second * 3)
}
