package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("main start")
	ch := make(chan string)
	go func() {
		for _, s := range os.Args[1:] {
			ch <- s
		}

		close(ch)
	}()

	for s := range ch {
		fmt.Printf("main----%s", s)
	}
}
