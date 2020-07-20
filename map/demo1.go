package main

import "fmt"


func main() {
	var m1  = map[int]interface{}{1:2, 2:4}
	var m2  = map[int]interface{}{1:"hello", 2:"world"}
	fmt.Println(m2)
	fmt.Println(m1)

	fmt.Println(m1)
	fmt.Println(m2)
}
