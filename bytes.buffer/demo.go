package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	buf := bytes.NewBuffer([]byte("hello!"))
	buf.WriteString("小花猫")
	fmt.Println(buf.String())

	buf1 := bytes.NewBuffer([]byte{})
	s := []byte("小黑猫")
	buf1.Write(s)
	fmt.Println(buf1.String())

	buf3 := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o'})
	var b byte = '?'
	buf3.WriteByte(b)
	fmt.Println(buf3.String())

	buf5 := bytes.NewBuffer([]byte{})
	var r rune = '小'
	buf5.WriteRune(r)
	fmt.Println(buf5.String())

	file, _ := os.Open("./text.txt")
	buf4 := bytes.NewBufferString("hello")
	buf4.WriteTo(file)

}
