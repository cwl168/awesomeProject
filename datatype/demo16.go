package main

import "fmt"

func main() {
	bytes := []byte{''}
	fmt.Print(bytes)
	/*var (
		c1 byte = 'a'
		c2 byte = '新'
		c3 rune = '新'
	)
	fmt.Println(c1, c2, c3)*/
	fmt.Printf("0x%x, %d\n", '', '')
	fmt.Println(0x81 == '', '\u0081' == '', 129 == '')

	a := [...]int{1: 100, 2: 200, 3: 200}
	b := a[0:]
	fmt.Println(b)
}
