package main

import "fmt"

type MyInt int

func (a MyInt) Add(b MyInt) MyInt { //面向对象
	return a + b
}
func main() {
	var a MyInt = 1
	var b MyInt = 2
	fmt.Println("a.Add(b) = ", a.Add(b))
}
