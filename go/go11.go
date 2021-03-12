package main

import (
	"fmt"
	//"sync"
	"time"
)

//var m sync.Mutex  //Go实现单机并发缓存 http://books.studygolang.com/gopl-zh/ch9/ch9-07.html
var set = make(map[int]bool, 0)

func printOnce(num int) {
	//m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	//m.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
