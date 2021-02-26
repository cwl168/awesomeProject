package main

import (
	"fmt"
	"sync"
)

// 定义Once对象
var once sync.Once

func main() {
	onceBody := func() {
		fmt.Println("我只跑一次。")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		// 并发执行10次
		go func() {
			// 调用Once对象的Do函数执行onceBody函数
			// 注意：一个Once对象只能执行一次，无论调用多少次Do函数。
			once.Do(onceBody)

			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		<-done
	}
}
