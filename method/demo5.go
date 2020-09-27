package main

import (
	"fmt"
)

//方法表达式：也即“方法对象赋值给变量”
//两种使用方式：
//1）隐式调用, struct实例获取方法对象---->方法值

//2)显示调用, struct类型获取方法对象, 须要传递struct实例对象作为参数。---->方法表达式

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) PrintInfoPointer() {
	fmt.Printf("%p, %v\n", p, p)
}

func (p Person) PrintInfoValue() {
	fmt.Printf("%p, %v\n", &p, p)
}

const s = "Go101.org"

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func main() {
	println(len(s), len(s[:]), a, b)
	//直接调用
	p := Person{"ck_god", 'm', 18}
	p.PrintInfoPointer()
	fmt.Println("---------------\n")

	//方法表达式
	pFunc1 := (*Person).PrintInfoPointer
	pFunc1(&p)

	pFunc2 := Person.PrintInfoValue
	pFunc2(p)
	fmt.Println("---------------\n")

	//方法值
	pFunc3 := p.PrintInfoPointer
	pFunc3()

	pFunc4 := p.PrintInfoValue
	pFunc4()
	fmt.Println("---------------\n")

	//备注：pFunc2和pFunc4的内存地址是不一样的；pFunc1和pFunc3的内存地址是一致的
}
