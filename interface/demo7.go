package main

import "fmt"

type animal interface {
	printInfo()
}
type cat int

//值接收者实现animal接口
func (c *cat) printInfo() {
	fmt.Println("a cat")
}

//需要一个animal接口作为参数
func invoke(a animal) {
	a.printInfo()
}
func main() {
	//参数c是一个变量，编译器会隐式获取它的地址
	var c cat
	//值作为参数传递
	invoke(c)
}

//cannot use c (type cat) as type animal in argument to invoke:
//cat does not implement animal (printInfo method has pointer receiver)
//cat没有实现animal接口，因为printInfo方法有一个指针接收者,所以cat类型的值c不能作为接口类型animal传参使用，所以这就是为啥书里面说T类型的值不拥有所有*T指针的方法。
