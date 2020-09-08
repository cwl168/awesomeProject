package main

import "fmt"

func zero1(ptr *[32]byte) {
	*ptr = [32]byte{}
}
func main() {
	var te [32]byte
	zero1(&te)
	for i, v := range te {
		fmt.Println(i, v)
	}
	x := 2
	y := 64
	fmt.Println(x << 1)
	fmt.Println(y >> 5)
}
