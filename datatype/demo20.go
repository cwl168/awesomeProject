package main

import (
	"fmt"
)

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	a := rotate(testArr[:], 2) //[0, 1, 2, 3, 4]变为[2 3 4 0 1]
	fmt.Println(a)
}
func rotate(s []int, r int) []int {
	lens := len(s)
	//创建一个空的指定长度的slice
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}
