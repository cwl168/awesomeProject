package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++{
		select {
		case ch <- i:
		case x := <- ch:
			fmt.Println(x)
		}
	}
}