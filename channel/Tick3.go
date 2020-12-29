package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for n := 10; n > 0; n-- {
		// 有意思的是，一旦 tick 或者 abort 这两个 channel 有任何一个可读了，
		// 那么就会进入相应的 case 分支
		select {
		case <-tick:
			fmt.Println(n)
		case <-abort:
			fmt.Println("abort")
			return
		}
	}
	fmt.Println("launch")
}
