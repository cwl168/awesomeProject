package main

import (
	"fmt"
	"regexp"
	"strings"
)

func add1(r rune) rune { return r + 1 }
func main() {
	fmt.Println(strings.Map(add1, "HAL-9000"))

	s1 := "hello,world"
	s2 := s1[:5]
	//fmt.Println(cap(s1))   //panic
	fmt.Println(len(s1)) //11
	//fmt.Println(cap(s2))   //panic
	fmt.Println(len(s2))   //5
	fmt.Printf("%T\n", s2) //string

	/*a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := a[:4]
	s2 := s[:7]
	fmt.Println(s, s2)    //[1 2 3 4] [1 2 3 4 5 6 7]
	fmt.Println(cap(s))   //10
	fmt.Println(s[7])     //panic:index out of range
	fmt.Printf("%T\n", s) //[]int*/
	sourceStr := "http://top.baidu.com/?fr=mhd_card"
	if ok, _ := regexp.MatchString(`.*javascript+`, sourceStr); ok {
		fmt.Println("ddddd")
	}
}
