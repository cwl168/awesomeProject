package main

import "fmt"
const a = iota
const b = iota
const c = iota
const (
	d = iota
	e = iota
	f
	_
	g
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
