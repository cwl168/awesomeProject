package main

import (
	"fmt"
)

//只能对命名类型添加方法
func (c *int) Write(p []byte) (int, error) {
	*c += len(p)
	return len(p), nil
}
func main() {
	var c int = 0
	var name = "Dolly"
	d, _ := fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	fmt.Println(d)
}
