package main

import (
	"fmt"
	"learnDemo/method/person"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
func main() {
	var q = person.Test{"cwl"}
	fmt.Println(q)
	var p = person.NewPerson("cwl")
	(*p).Setage(31)
	fmt.Println(*p)
	fmt.Printf("%T\n", q)
	fmt.Println(fmt.Sprintf("%T\n", q))

	var c ByteCounter = 0
	var name = "Dolly"
	d, _ := fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	fmt.Println(d)
}
