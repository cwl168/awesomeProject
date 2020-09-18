package main

import (
	"fmt"
	"unicode"
)

// 示例：判断大小写
func main() {
	s := "Hello 世界！"

	for _, r := range s {
		fmt.Printf("%c", unicode.ToUpper(r))
	} // HELLO 世界！
	for _, r := range s {
		fmt.Printf("%c", unicode.ToLower(r))
	} // hello 世界！
	for _, r := range s {
		fmt.Printf("%c", unicode.ToTitle(r))
	} // HELLO 世界！

	for _, r := range s {
		fmt.Printf("%c", unicode.To(unicode.UpperCase, r))
	} // HELLO 世界！
	for _, r := range s {
		fmt.Printf("%c", unicode.To(unicode.LowerCase, r))
	} // hello 世界！
	for _, r := range s {
		fmt.Printf("%c", unicode.To(unicode.TitleCase, r))
	} // HELLO 世界！
}
