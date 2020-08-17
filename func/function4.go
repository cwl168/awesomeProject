package main

import "fmt"

func test() int {
	var x int
	x++
	return x
}
func test2() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
func main() {
	//Test( )保存在栈中，函数结束，栈释放内存
	/*fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())*/
	//虽然Test()已经返回了，但是返回的值：func（）还在全局变量中使用,三次调用 f（），因此返回值会保存在堆上，即使栈释放了内存资源，但func（）保存在堆中，数据不会释放。
	f := test2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
