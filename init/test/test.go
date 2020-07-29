package test

import (
	"fmt"
	"learnDemo/init/test2"
)

func init() {
	fmt.Println("1st init print in test package")
}

func TestFunc() {
	test2.TestFunc2()
}
