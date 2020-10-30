package main

import (
	"fmt"
	"runtime"
	"time"
)

func sayHello() {
	for {
		fmt.Println("Hello gorotine")
		time.Sleep(time.Second)
	}
}

func main() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	go sayHello()
	fmt.Println("Hello main")
}
