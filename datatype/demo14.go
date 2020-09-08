package main

import "fmt"

const (
	deadbeef = 0xdeadbeef
	a        = uint32(deadbeef)
	//d = int32(deadbeef)
)

func main() {
	fmt.Println(a)
	i := 0
	r := '\000'
	f := 0.0
	c := 0i
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", c)

	//var a [3]int
	type Currency int
	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	b := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	for i, v := range b {
		fmt.Printf("%d %s\n", i, v)
	}

	t := [...]int{99: -1}
	for i, v := range t {
		fmt.Printf("%d %d\n", i, v)
	}
}
