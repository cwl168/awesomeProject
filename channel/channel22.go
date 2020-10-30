package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer has expired.")
	}()
	timer.Stop()
	//timer.Reset(2  * time.Second)
	time.Sleep(5 * time.Second)
}
