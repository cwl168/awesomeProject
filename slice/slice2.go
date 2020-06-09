package main

import "fmt"

func main() {
	str := [4]string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}
	fmt.Println(str)

	a := str[0:2]
	b := str[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(str)
}