package main

import "fmt"

func main() {
	medals := []string{}
	//如果len函数返回一个无符号数，那么i也将是无符号的 uint类型
	i := uint(len(medals)) - 1
	fmt.Println(i)
	fmt.Printf("%T", i)
}
