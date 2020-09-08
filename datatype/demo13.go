package main

import (
	"fmt"
)

func main() {
	var f float64 = 212
	//fmt.Println((f - 32) * 5 / 9)
	//fmt.Printf("%T\n", (f-32)*5/9)
	//
	//fmt.Println(5 / 9 * (f - 32))
	//fmt.Printf("%T\n", 5/9*(f-32))

	fmt.Println(5.0 / 9.0 * (f - 32))
	fmt.Printf("%T\n", 5.0/9.0*(f-32))
	fmt.Printf("%v=>%[1]T\n", 5.0/9.0*(f-32))

	//fmt.Printf("%T\n",5 / 9)
	//fmt.Println(5.0 / 9.0 * (f - 32))
	//
	//var f1 float64 = 3 + 0i
	//f1 = 2
	//fmt.Printf("%T\n", f1)
}
