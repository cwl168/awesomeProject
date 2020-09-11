package main

import (
	"fmt"
)

func main() {
	bytes := []byte("hello world")
	copy(bytes, "ha ha")
	fmt.Println(bytes)
	b := []string{"tao", "taoshihan", "taoshihan", "shi", "shi", "han"}
	fmt.Println(emptyString(b))
}
func emptyString(s []string) []string {
	k := 0
	for _, w := range s {
		if s[k] == w {
			continue
		}
		k++
		s[k] = w
	}
	return s[:k+1]
}
