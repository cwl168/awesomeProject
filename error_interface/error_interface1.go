package main

import (
	"fmt"
	"runtime"
)

type stack []uintptr
type errorString struct {
	s string
	*stack
}

type withMessage struct {
	cause error
	msg   string
}

func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}
func (fe *errorString) Error() string {
	return fe.s + "3333"
}
func New(text string) error {
	return &errorString{
		s:     text,
		stack: callers(),
	}
}
func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

//只是模拟一个错误
func openFile() ([]byte, error) {
	return nil, &fileError{"文件错误，自定义"}
}
func main() {
	conent, err := openFile()
	fmt.Printf("%T\n", err)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(conent))
	}
}
