package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
a[i]++  中间没有机会进行协程切换，会被一个协程所抢占，没法交出控制权，始终在这个协程里面
fmt.Printf  io操作会有协程之间的切换，会交出控制权
i等于10，导致运行中协程的a[0]++,a[1]++,a[2]++ ... 都变成a[10]++ ，导致越界

*/
func main() {
	var a [11]int
	for i := 0; i < 10; i++ {
		//i := i
		go func() {
			for {
				a[i]++
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(a)

}
