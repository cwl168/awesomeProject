package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	aa := []rune("中文")

	fmt.Println(len(aa))
	s := "abc"
	b := []byte(s)
	b[1] = 'f' //变量b被修改的情况下，原始的s字符串也不会改变
	s2 := string(b)
	//s2[1] = 'v'
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(s2)

	f := strings.Fields("ac fv gf")
	fmt.Println(f)

	g := bytes.Contains(b, []byte{97})
	fmt.Println(g)
}
