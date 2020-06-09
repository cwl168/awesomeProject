package main

import "fmt"

func main() {
	var numbers = make([]int,3,5)

	var numbers3 = [5]int{1, 2, 3, 4, 5}

	var slice1 = numbers3[1:4] //2 3 4

	var slice2 = slice1[1:3] //3 4
	for i,x:= range slice2 {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}

	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}