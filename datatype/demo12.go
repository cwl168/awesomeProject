package main

import (
	"fmt"
	"math"
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
	const (
		f = 1 << iota
		g
		h
		i
		k
	)
	fmt.Println(f, g, h, i, k)

	var x1 float32 = math.Pi
	var x2 float64 = math.Pi
	var x3 float32 = float32(x2)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)

	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776 (exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424 (exceeds 1 << 64)
		YiB // 1208925819614629174706176
	)
	fmt.Println(YiB / ZiB)

	var t float64 = 212
	fmt.Println((t - 32) * 5 / 9)
}
