package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan int)
	wg.Add(1)
	go produce(done, &wg)
	wg.Add(1)
	go consume(done, &wg)
	wg.Wait()
}
func produce(done chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		done <- i
		fmt.Printf("produce num %d\n", i)
	}
	close(done)
	wg.Done()
}
func consume(done <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < 8; i++ {
		fmt.Printf("consume num %d\n", <-done)
	}
	wg.Done()
}
