package main

import (
	"fmt"
	"sync"
)

func add(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count = *count + 1
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	count := 0
	wg.Add(3)
	go add(&count, &wg)
	go add(&count, &wg)
	go add(&count, &wg)

	wg.Wait()
	//可运行多次的结果，都不相同
	fmt.Println("count 的值为：", count)
}
