package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

//go run -gcflags "-m -l" ./func/function6.go
func main() {
	x := foo()
	fmt.Println(*x)
}
