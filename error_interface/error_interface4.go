package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("whoops")
	fmt.Printf("%+v", err) //调用堆栈

	fmt.Println()

	cause := errors.New("whoops")
	err1 := errors.WithMessage(cause, "oh noes")
	fmt.Println(err1)

	fmt.Println()

	cause1 := errors.New("whoops")
	err2 := errors.Wrap(cause1, "oh noes")
	fmt.Println(err2)

	//fmt.Println(err)

}
