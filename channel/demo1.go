package main

import "fmt"

func main() {
	ch := make(chan int)
	n := 10
	go foo(n, ch)

	for i := 0; i < n; i++ {
		t := <-ch
		if t == 2 {
			return
		}
		fmt.Println(t)
	}
}
func foo(n int, in chan<- int) {
	for i := 0; i < n; i++ {
		in <- i
	}
}
