package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	fmt.Println("it is the main goroutine")
	time.Sleep(time.Second * 1)
}

func test() {
	fmt.Println("it is a new goroutine")
}