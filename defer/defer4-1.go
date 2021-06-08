package main

import (
	"fmt"
)

func main() {
	fmt.Println("return:", b()) //  打印结果为  return:  2
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) //  打印结果为  defer:  2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) //  打印结果为  defer:  1
	}()
	return
}

//defer1: 1
//defer2: 2
//return: 2
