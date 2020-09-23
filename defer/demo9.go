package main

import "fmt"

func double(x int) int {
	return x + x
}
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

//先对返回值进行赋值
//执行 defer 语句
//执行函数返回
// return会将返回值先保存起来，
//对于无名返回值来说，保存在一个临时对象中，defer是看不到这个临时对象的；
// 而对于有名返回值来说，就保存在已命名的变量中。
func main() {
	fmt.Println(triple(4)) //12
}
