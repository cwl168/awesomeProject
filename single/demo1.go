package main

import (
	"fmt"
	"sync"
)

type example struct {
	name string
}

var instance *example

func GetExample() *example {

	// 存在线程安全问题，高并发时有可能创建多个对象
	if instance == nil {
		instance = new(example)
	}
	return instance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			obj := GetExample()
			fmt.Printf("%d,%p\n", i, obj)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
