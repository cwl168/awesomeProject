package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func do_stuff2(x int) int { // 一个比较耗时的事情，比如计算
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模拟计算
	return 100 - x                                              // 假如100-x是一个很费时的计算
}

var ch = make(chan int)

func fanIn2(nums ...int) {
	defer close(ch)
	for _, num := range nums {
		ch <- do_stuff2(num)
	}
}
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	go fanIn2(1, 2, 3)
	for num := range ch {
		fmt.Println(num)
	}
	//for i:=0; i < 3; i++ {
	//	fmt.Println(<-result)
	//}
}
