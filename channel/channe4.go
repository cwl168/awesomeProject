package main

import "fmt"


var ch8 = make(chan int, 6)

func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	close(ch8)
}
func main() {
	go mm1()
	for i := 0; i < 100; i++ {
		fmt.Print(<-ch8, "\t")
	}
}