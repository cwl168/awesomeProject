package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println("End:", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
