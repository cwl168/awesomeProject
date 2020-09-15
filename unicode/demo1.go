package main

import (
	"fmt"
	"unicode"
)

// 示例：判断汉字
func main() {
	for _, r := range "Hello 世界！" {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c", r) // 世界
		}
	}
}
