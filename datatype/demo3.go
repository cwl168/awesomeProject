package main

import "fmt"

func main() {
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(len(medals))
	for i := uint(len(medals)) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) //  越界
		//fmt.Println(i) // 死循环
	}
}
