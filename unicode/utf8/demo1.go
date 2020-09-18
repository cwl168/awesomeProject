package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := make([]byte, utf8.UTFMax)

	n := utf8.EncodeRune(b, '好') //好UNICODE字符由2个字节表示为597D，UTF-8需要3个字节 E5A5BD， utf-8編碼为E5A5BD，需要三个字节表示
	fmt.Printf("%v：%v\n", b, n)  // [229 165 189 0]：3    E5 229    A5 165  BD 189

	s := "大家好"
	for i := 0; i < len(s); {
		r, n := utf8.DecodeRuneInString(s[i:])
		fmt.Println(s[i:])

		fmt.Printf("%c：%v\n", r, n) // 大：3   家：3   好：3
		i += n
	}
	fmt.Println()
}
