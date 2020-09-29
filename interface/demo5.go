package main

import (
	"fmt"
)

type a struct {
	v int
}

func (i *a) test() {
	i.v = 666
}

func main() {
	t := a{1}
	fmt.Println(t) //{1}
	t.test()
	fmt.Println(t) //{666}

	p := &a{1}
	fmt.Println(p.v) //1
	p.test()
	fmt.Println(p.v) //666

	var b a
	b.test()
	fmt.Println(b.v)
	//var _ = a{1}.test() //error
	/*
		  # command-line-arguments
		./a.go:27:15: cannot call pointer method on a literal
		./a.go:27:15: cannot take the address of a literal
		./a.go:27:20: a literal.test() used as value
	*/
}
