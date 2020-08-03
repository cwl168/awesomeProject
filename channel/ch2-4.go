package main

import (
	"fmt"
	"runtime"
	"time"
)

//goroutine 泄露

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Printf("goroutine num %d", runtime.NumGoroutine())
	}()
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
	/*for i := 0; i < 3; i++ {
		fmt.Printf(" %d\n", <-ch)
	}*/
}
