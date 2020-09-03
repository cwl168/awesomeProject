package main

import (
	"fmt"
)

//一个在函数体词法域，一个在for隐式 的初始化词法域，一个在for循环体词法域
func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
}

// "HELLO" (one letter per iteration)
