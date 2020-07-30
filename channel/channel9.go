package main

import "fmt"

func hello(pipline chan string) {
	pipline <- "hello world"

}
func main() {
	/*pipline := make(chan string)
	pipline <- "hello world"  写端阻塞
	fmt.Println(<-pipline)*/

	/*pipline := make(chan string)
	fmt.Println(<-pipline)	读端阻塞
	pipline <- "hello world" */

	pipline := make(chan string)
	go hello(pipline)
	fmt.Println(<-pipline)

}
