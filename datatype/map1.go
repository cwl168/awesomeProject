package main

import "fmt"

func main() {
	ages := make(map[int]int)
	ages[1] = 31
	ages[2] = 34
	for name, age := range ages {
		fmt.Printf("%d\t%d\n", name, age)
	}
	fmt.Println(getKeys1(ages))
}
func getKeys1(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		fmt.Println(k)
		j++
	}
	return keys
}
func getKeys2(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 初始化默认
func getKeys3(m map[int]int) []int {
	// 注意：由于数组默认长度为0，后面append时，需要重新申请内存和拷贝，所以效率较低
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
