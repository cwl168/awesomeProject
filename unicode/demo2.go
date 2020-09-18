package main

import (
	"fmt"
	"unicode"
)

// 示例：判断大小写
func main() {
	for _, r := range "Hello ＡＢＣ！" {
		// 判断字符是否为大写
		if unicode.IsUpper(r) {
			fmt.Printf("%c", r) // HＡＢＣ
		}
	}
	for _, r := range "Hello ａｂｃ！" {
		// 判断字符是否为小写
		if unicode.IsLower(r) {
			fmt.Printf("%c", r) // elloａｂｃ
		}
	}
	for _, r := range "Hello ᾏᾟᾯ！" {
		// 判断字符是否为标题
		if unicode.IsTitle(r) {
			fmt.Printf("%c", r) // ᾏᾟᾯ
		}
	}
}
