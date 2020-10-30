package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	total := 0
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("total: ", total)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	var mutex sync.Mutex
	for i := 0; i < 2; i++ {
		go func() {
			mutex.Lock()
			total += 1
		}()
	}
}
