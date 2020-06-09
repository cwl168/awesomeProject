package main

import "fmt"

func main() {
	pipline := make(chan string)
	go func() {
		pipline <- "hello world"
		pipline <- "hello China"
		close(pipline)
	}()
	for data := range pipline{
		fmt.Println(data)
	}
}