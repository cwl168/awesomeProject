package main

import (
	"fmt"
)

func main() {
	/*ch := make(chan string)
	go func() {
		ch <- "send"
	}()
	fmt.Println(<-ch)*/

	/*ch := make(chan string)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- "send"*/

	/*chs := make(chan string, 2)
	chs <- "first"
	chs <- "second"
	close(chs)

	for ch := range chs {
		fmt.Println(ch)
	}*/

	ch := make(chan string)
	ch1 := make(chan string)
	ch1 <- "channelValue"
	fmt.Println(<-ch)

	/*var done = make(chan struct{})
	select {
	case <-done:
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}
	close(done)*/
}
