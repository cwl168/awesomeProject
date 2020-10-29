package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)

		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}
