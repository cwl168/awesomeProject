package main

import (
	"fmt"
)

func main() {
	fmt.Println("return:", a()) //  打印结果为  return:  0
}

//先对返回值进行赋值
//执行 defer 语句
//执行函数返回

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) //  打印结果为  defer:  2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) //  打印结果为  defer:  1
	}()
	return i
}

//defer1: 1
//defer2: 2
//return: 0
