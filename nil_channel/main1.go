package main

import (
	"fmt"
)

//https://medium.com/@ashishstiwari/dont-simply-run-forever-loop-for-1594464040b1
func main() {
	go anotherGoroutine()
	forever()
}
func forever() {
	for {
	}
}
func anotherGoroutine() {
	i := 0
	for {
		i++
		fmt.Print(i)
		fmt.Print("\n")
	}
}
