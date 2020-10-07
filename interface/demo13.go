package main

import "fmt"

func main() {
	var a, b, c int
	for i := 0; i < 2; i++ {
		fmt.Scanf("%d,%d,%d", &a, &b, &c)
		fmt.Println(a, b, c)
	}
}
