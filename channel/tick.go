package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Tick(3 * time.Second)

	for {
		select {
		case <-t1:
			fmt.Println("t1定时器")
		}

	}
}
