package main

import (
	"fmt"
	"sync"
)

var x, y int

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// goroutine A
	go func() {
		x = 1
		fmt.Printf("y: %d ", y)
		wg.Done()
	}()

	// goroutine B
	go func() {
		y = 1
		fmt.Printf("x: %d ", x)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("end")
}
