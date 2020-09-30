package main

import "fmt"

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

//在 main 函数中，调用调用 sayHello() 函数时，传入了 golang, php 对象，它们并没有显式地声明实现了 IGreeting 类型，只是实现了接口所规定的 sayHello() 函数。实际上，编译器在调用 sayHello() 函数时，会隐式地将 golang, php 对象转换成 IGreeting 类型，这也是静态语言的类型检查功能。
func main() {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}
