package main

import (
	"fmt"
)

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array[:3])
	s := make([]int, 0, 1)
	fmt.Printf("len:%d,cap:%d\n", len(s), cap(s))
	c := cap(s)
	for i := 0; i < 2000; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
	//fmt.Println(s)
}
