package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	/*x := <- ch
	fmt.Println(x)*/
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
