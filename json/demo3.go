package main

import "fmt"

type student struct {
	id   int32
	name string
}

func main() {
	a := &student{id: 1, name: "xiaoming"}

	fmt.Printf("a=%v	\n", a)
	fmt.Printf("a=%+v	\n", a)
	fmt.Printf("a=%#v	\n", a)
}
