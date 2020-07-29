package main

import (
	"fmt"
	"learnDemo/init/test"
)

func init() {
	fmt.Println("1st init print in main package")
}

func main() {
	test.TestFunc()
}
