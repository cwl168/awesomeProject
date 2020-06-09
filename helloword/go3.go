package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go printHelloWorld3(i, ch)
	}

	//fatal error: all goroutines are asleep - deadlock!  在接受的时候，一定得有发送，没发送会deadlock
	for  {
		msg := <-ch
		fmt.Println(msg)
	}
}
func printHelloWorld3(i int, ch chan string) {
		ch <- fmt.Sprintf("hello world %d\n", i,)

}
