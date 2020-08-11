package main

import "fmt"

func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}
func void() {

}

//命令行来分析变量逃逸   go run -gcflags "-m -l" ./func/function5.go
func main() {
	// 声明a变量并打印
	var a int
	// 调用void函数

	void()
	fmt.Println(a, dummy(0))

}
