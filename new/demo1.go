package main

import "fmt"

func main() {
	m := make(map[int]int, 10)
	p := new(bool)
	fmt.Println(*p)
	_, ok := m[1]
	v := m[2]
	fmt.Println(ok)
	fmt.Println(v)
	fmt.Println(gcd(15, 20))
}

//最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
