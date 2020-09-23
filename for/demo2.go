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
 go build -gcflags '-m -l' for/demo2.go

[7558691 7537657 7554229 7546357 7533731 7573485 7552624 7561120 7555174 7522028 0]   //0-10
[0       0 			0 		2 		8 		4 		6 		4 		3 		82 	 36416829]  //0-10
*/
func main() {
	var a [11]int
	for i := 0; i < 10; i++ {
		i := i
		fmt.Println(i)
		go func() {
			for {
				a[i]++

				runtime.Gosched()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
