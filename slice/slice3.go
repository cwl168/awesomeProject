package main

import "fmt"

func main() {
	//slice长度为101,  b[0]= 1 .... b[3]= 4 b[4]=0 ... b[99]=0, b[100]=1
	c := []uint32{1, 2, 3, 4, 100: 1}
	fmt.Printf("%+v,%d\n", c, len(c))

	a := [...]int{1, 2, 3, 10: 1}
	b := []int{1, 2, 3, 10: 1}
	fmt.Println(a, b)
}
