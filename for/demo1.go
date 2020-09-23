package main

import "fmt"

var slice []func()

//捕获迭代变量
func main() {
	sli := []int{1, 2, 3, 4, 5}
	for _, v := range sli {
		fmt.Println(&v)
		slice = append(slice, func() {
			fmt.Println(v * v) // 直接打印结果
		})
	}

	for _, val := range slice {
		val()
	}
}
