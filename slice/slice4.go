package main

import "fmt"

func main() {
	var ints []int = make([]int, 2, 5)
	ints = append(ints, 1)
	fmt.Println(ints)
	ints = append(ints, 2)
	fmt.Println(ints)
	fmt.Println(ints[3])

	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3:3]
	newSlice[1] = 45
	newSlice = append(newSlice, 46)
	newSlice = append(newSlice, 47)
	newSlice = append(newSlice, 48)
	newSlice = append(newSlice, 49)
	newSlice = append(newSlice, 50)
	fmt.Println(newSlice)
	fmt.Println(slice)

	s := []string{"a", "b", "c", "d", "e", "f"}
	s1 := s[2:3:4]
	s1 = append(s1, "x")
	s1 = append(s1, "y")
	fmt.Println(s)
	fmt.Println(s1, len(s1), cap(s1))

	s3 := []int{1, 2, 3}
	s4 := []int{4, 5}
	s3 = append(s3, s4...)
	fmt.Println(s3)

	list := new([]int)

	*list = append(*list, 1)

	fmt.Println(*list)

	p1 := new([]int)
	fmt.Println(*p1)

	p2 := new(int)
	fmt.Println(*p2)

	p3 := new(string)
	fmt.Println(*p3)

	p4 := new(bool)
	fmt.Println(*p4)

}
