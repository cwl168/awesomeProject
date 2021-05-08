package main

import "fmt"

func main() {
	slice := make([]int, 2, 4)
	slice = append(slice, 4)
	rz := appendInt(slice, 3)
	fmt.Println(rz)
	rz = appendInt(rz, 4)
	fmt.Println(rz)

	rz = appendInt(rz, 5)
	fmt.Println(rz)
	rz = appendInt(rz, 6)
	fmt.Println(rz)
	rz = appendInt(rz, 7)
	fmt.Println(rz)
	rz = appendInt(rz, 8)
	fmt.Println(rz)
	rz = appendInt(rz, 9)
	fmt.Println(rz)
	rz = appendInt(rz, 10)
	fmt.Println(rz)

}
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
