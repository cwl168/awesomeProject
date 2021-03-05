package main

import (
	"fmt"
	"time"
)

var dummychan chan int // default chan, has nil value
func poorFoo(done chan bool, ch <-chan int) {
	select {
	case x := <-ch: // forever blocked on a nil channel
		fmt.Printf(" Received %d", x)
	case <-done: // to listen to signal on this channel
		return
	case <-time.After(1 * time.Second):
		fmt.Println(" Timed out, bye")
		return
	}
}

func main() {
	donechan := make(chan bool)
	fmt.Println("In maincaller")
	go poorFoo(donechan, dummychan)
	// do some activity as usual and run as long as you need
	time.Sleep(time.Second * 10)

	close(donechan) //  signal while exiting
}
