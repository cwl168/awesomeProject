package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Tick 返回一个 <-chan Time 类型的 channel
	// 这个 channel 每隔固定时间，吐出一个时间。这里我设置的 1 秒。
	tick := time.Tick(1 * time.Second)
	for n := 10; n > 0; n-- {
		fmt.Println(n)
		fmt.Println(<-tick)
	}
	fmt.Println("launch")
}
