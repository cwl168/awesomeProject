package main

import (
	"errors"
	"fmt"
	"syscall"
)

/*
errors包中的代码
package errors
//定义了接口
type error interface {
    Error() string
}
//大写字母开头的方法,可以导出
//返回了errorStrig类型
func New(text string) error {
	return &errorString{text}
}
//定义类型
type errorString struct {
	s string
}
//类型实现接口的方法
func (e *errorString) Error() string {
	return e.s
}
*/
func main() {
	//返回false,这俩是不相同的实例
	fmt.Println(errors.New("tsh error") == errors.New("tsh error"))

	//fmt.Errorf进行了包装
	fmt.Println(fmt.Errorf("我是 %s 错误", "tsh"))

	//类似实现了error接口
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
