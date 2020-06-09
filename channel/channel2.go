package main

import (
	"fmt"
	"time"
)

var ch10 = make(chan int, 6)

func mm3() {
	for i := 0; i < 10; i++ {
		ch10 <- 8 * i
	}
	close(ch10)
}
func mm4() {
	for {
		for data:=range ch10{
			fmt.Print(data,"\t")
		}
	}
}
func main() {
	go mm3()
	go mm4()
	time.Sleep(time.Second)
	/*for{
		runtime.GC()
	}*/
}
