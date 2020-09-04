package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "∂he" +
		"llo, \nworld\n" //∂ 三个字节
	fmt.Println(s)
	t := `∂hel
lo, \nworld\n` //∂ 三个字节
	fmt.Println(t)
	var g byte = 0
	fmt.Println(g)

	a := "你好，china" //utf-8编码下占3个字节 3*3+5   中文字符在unicode下占2个字节
	fmt.Println(len(a))
	fmt.Println(len([]rune(a)))
	fmt.Println(utf8.RuneCountInString(a))

	//var str1 = "go算法"
	//fmt.Println(str1[:4])
	//
	//var str = "go算法"
	//for k, v := range str {
	//	fmt.Printf("v type: %T\n", v)
	//	fmt.Println(v,k)
	//}
	//str5:="hello 世界"
	//for i:=0;i<len(str5);i++{
	//	fmt.Printf("%c",str5[i]) // hello ä¸ç
	//}
	//
	//str2:="hello 世界"
	//for _,r:=range str2{
	//	fmt.Printf("%c\n",r) // hello 世界
	//}
	//
	//var b byte = '1'
	//var c rune = '中'
	//fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数", unsafe.Sizeof(b), unsafe.Sizeof(c))
}
