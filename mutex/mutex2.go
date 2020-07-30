package main

import (
	"fmt"
	"sync"
)

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	count := 0
	wg.Add(3)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)

	wg.Wait()
	fmt.Println("count 的值为：", count)
}
