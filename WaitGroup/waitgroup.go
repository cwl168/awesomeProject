package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}