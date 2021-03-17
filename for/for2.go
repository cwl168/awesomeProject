package main

import "fmt"

//退出本层循环 可以选择中断哪一个循环
func main() {
	for j := 0; j < 5; j++ {
	loop:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break loop
			}
			fmt.Println(i)
		}
	}
}
