package main

import "fmt"

var ch8 = make(chan int)

func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	close(ch8)
}
func main() {
	go mm1()
	for i := 0; i < 11; i++ {
		fmt.Print(<-ch8, "\t")
	}
	//for num := range ch8 {
	//	fmt.Println(num)
	//}
}
