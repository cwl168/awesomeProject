//所谓高阶函数就是将一个或多个其他函数作为自身函数的入参，并在函数体内调用它们
package main

import "fmt"

//自定义一个函数类型
type F func(int) int

//实现高阶函数
func HigherFunction(v int, f F) int {
	return f(v)
}

//调用高阶函数
func main() {
	//声明函数1
	var f1 F
	f1 = func(i int) int {
		return i * 2
	}
	rs1 := HigherFunction(2, f1)
	fmt.Println("执行高阶函数1结果:", rs1)

	//声明函数2
	var f2 F
	f2 = func(i int) int {
		return i * 20
	}
	rs2 := HigherFunction(2, f2)
	fmt.Println("执行高阶函数2结果:", rs2)

}
