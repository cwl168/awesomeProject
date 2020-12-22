package main

import (
	"fmt"
)

func main() {
	requests := []int{12, 2, 3, 41, 5, 6, 1, 12, 3, 4, 2, 31}
	for n := range requests {
		fmt.Println(n)
		go run(n) //开启多个协程
	}

	/*for {
		select {}
	}*/
}

func run(num int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	//模拟请求错误
	if num%5 == 0 {
		panic("请求出错")
	}
	go myPrint(num)
}

//Go的Web服务也是一样，默认的recover机制只能捕获一层，如果你在这个请求的处理中又使用了其它协程，那么必须非常慎重，毕竟只要发生一个panic，整个Web服务就会挂掉。
func myPrint(num int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	if num%4 == 0 {
		panic("请求又出错了")
	}
	fmt.Printf("%d\n", num)
}
