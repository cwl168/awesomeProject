package main

import (
	"fmt"
	"learnDemo/import_cycle/b"
)

func main() {
	fmt.Println(b.Result())
}

// main 引入b包 ，Result方法调用 引入a 包，Calculate方法调用 引入 b包
