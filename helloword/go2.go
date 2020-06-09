package main

import (
	"fmt"
	"time"
)

/**
 5个 在 1毫秒 争着输出hello world
 */
func main() {
	for i := 0; i < 5; i++ {
		go printHelloWorld2(i)
	}
	time.Sleep(time.Millisecond)
}
func printHelloWorld2(i int) {
	for {
		fmt.Printf("hello world %d\n", i)
	}
}
