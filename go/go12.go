package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//fmt.Println(runtime.NumCPU()) //一个双核CPU，带有超线程技术，则会被认为是4个逻辑CPU
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)

	// 模拟 Goroutine 死循环
	go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("脑子进煎鱼了")

}
