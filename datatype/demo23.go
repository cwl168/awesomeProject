package main

import "fmt"

//用于将内容从一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中  [1 2 3]
	fmt.Println(slice2)
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置  [5 4 3 4 5]
	//fmt.Println(slice1)
	//http://www.cnblogs.com/osfipin/
}
