package main

import "fmt"

func main() {
	//s := "∂hello, world"  //∂ 三个字节
	//s := "hello, world"
	s := "hello,世界\377" //汉字  三个字节
	s += ", right foot"
	fmt.Println(len(s))
	//s[0] = 'L' //字符串的值是不可变的
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
}
