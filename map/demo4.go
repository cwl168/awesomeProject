package main

import (
	"fmt"
	"unicode/utf8"
)

type info struct {
	result int
}

func main() {
	myvar := 1
	var data info
	data.result = work()
	fmt.Println(data, myvar)

	/*m := make(map[string]int, 99)
	println(cap(m))     // error: invalid argument m1 (type map[string]int) for cap*/

	/*x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Println(v)    // 1 2 3
	}*/

	//Go 则会返回元素对应数据类型的零值，比如 nil、'' 、false 和 0，取值操作总有值返回，故不能通过取出来的值来判断 key 是不是在 map 中。
	/*x := map[string]string{"one": "2", "two": "", "three": "3"}
	if v := x["two"]; v == "" {
		fmt.Println("key two is no entry")    // 键 two 存不存在都会返回的空字符串
	}*/

	/*x := "text"
	xBytes := []byte(x)
	xBytes[0] = 'T'    // 注意此时的 T 是 rune 类型
	x = string(xBytes)
	fmt.Println(x)    // Text*/

	/*x := "text"
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x)    // 我ext*/

	char := "中"
	fmt.Println(len(char))
	fmt.Println(utf8.RuneCountInString(char))

	c := "é"
	fmt.Println(len(c))                    // 2
	fmt.Println(utf8.RuneCountInString(c)) //1

	x := []int{
		1,
		2,
	}
	fmt.Println(x)

	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
	i := 0
	//++i; error
	i++
	fmt.Println(i)
}
func work() int {
	return 3
}
