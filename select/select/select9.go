package main

import (
	"fmt"
)

func main() {
	var done = make(chan struct{})
	select {
	case <-done:
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}
	close(done)
}
