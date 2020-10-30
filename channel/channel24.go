package main

import (
	"fmt"
	"runtime"
	"time"
)

func UseTickerWrong() *time.Ticker {
	ticker := time.NewTicker(5 * time.Second)
	go func(ticker *time.Ticker) {
		for range ticker.C {
			fmt.Println("Ticker1....")
		}

		fmt.Println("Ticker1 Stop")
	}(ticker)
	return ticker
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	ticker1 := UseTickerWrong()
	time.Sleep(20 * time.Second)
	ticker1.Stop()
}
