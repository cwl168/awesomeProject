package main

import "fmt"

var name string = "go"


/**
无名返回值函数
return会将返回值先保存起来，对于无名返回值来说，
 保存在一个临时对象中，defer是看不到这个临时对象的；
 而对于有名返回值来说，就保存在已命名的变量中。
执行流程：
先对返回值进行赋值
执行 defer 语句
执行函数返回
 */
func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)//go
	//无名返回值
	return name  //go
}

func main() {
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)
}
/**
defer 是return 后才调用的。在执行 defer 前，myname 已经被赋值成 go
结果：
myfunc 函数里的name：go
main 函数里的name: python
main 函数里的myname:  go
 */