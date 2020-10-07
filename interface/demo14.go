package main

import "fmt"

type MyString string

type float float32

type Rect struct {
	Width  float
	Height float
}

func explain(i interface{}) {
	fmt.Printf("value give to explain function is of type:%T with value:%v\n", i, i)
}

func main() {
	var myString MyString
	myString = "hello"
	r := Rect{5.5, 4.5}
	explain(myString)
	explain(r)
}
