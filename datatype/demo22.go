package main

import (
	"fmt"
	"unicode"
)

func main() {
	d := []byte("a  bc")
	e := emptyString2(d)
	fmt.Println(string(e))
}
func emptyString2(s []byte) []byte {
	index := 0
	num := len(s)
	fmt.Println(s)
	for i := 0; i < num; i++ {
		index = i + 1
		num = len(s)
		if index >= num {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			//结合remove函数
			fmt.Println(s[i:], s[index:])
			copy(s[i:], s[index:])
			fmt.Println(s)
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
