package main

import (
	"fmt"
)

type T interface{}
type X string
type Y = string

func main() {

	var t T = "abc"
	var x X = "abc"
	var y Y = "abc"

	fmt.Println(t == x)
	fmt.Println(t == string(x))

	fmt.Println(t == y)
	fmt.Println(t == string(y))
}

// 输出什么?
