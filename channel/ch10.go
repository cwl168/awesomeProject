package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func worker(ch chan int) {
	defer wg.Done()
	for value := range ch {
		fmt.Println(value) // do something
	}
}

func main() {
	defer func() {
		fmt.Printf("goroutine num %d\n", runtime.NumGoroutine())
	}()
	ch := make(chan int)

	wg.Add(1)
	go worker(ch)

	for i := 0; i < 3; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
