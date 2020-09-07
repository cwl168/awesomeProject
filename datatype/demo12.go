package main

import (
	"fmt"
	"strconv"
	"time"
)

const pi = 3.14
const t = 1

//整数转为字符串
func main() {
	x := 256
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay) // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout) // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute)

	const (
		a = 1
		b
		c = 2
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}
