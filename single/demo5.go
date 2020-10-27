package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type example struct {
	name string
}

var instance *example

var mux sync.Mutex

var initialized uint32

func GetInstance() *example {

	// 一次判断即可返回
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mux.Lock()
	defer mux.Unlock()
	if initialized == 0 {
		instance = &example{}
		atomic.StoreUint32(&initialized, 1) // 原子装载
	}
	return instance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			obj := GetInstance()
			fmt.Printf("%d,%p\n", i, obj)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
