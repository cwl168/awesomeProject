package main

import "fmt"

func main() {
	aa := []rune("中文")

	fmt.Println(len(aa))
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(s2)
}
