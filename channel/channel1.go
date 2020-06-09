package main

import "fmt"

var ch9 = make(chan int, 6)

func mm2() {
	for i := 1; i < 10; i++ {
		ch9 <- 8 * i
	}
	close(ch9)

}
func main() {
	go mm2()
	for {
		if data, ok := <-ch9; ok {
			fmt.Print(data, "\t")
		} else {
			break
		}
	}
}
