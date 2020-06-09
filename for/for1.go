package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr01 := [...]int{1, 2, 3}
	arr01[2] = 4
	for k, v := range arr {
		fmt.Printf("%d=====%d\n", k, v)
	}
}
