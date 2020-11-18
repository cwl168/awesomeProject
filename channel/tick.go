package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Tick(3 * time.Second)
	c := []int{3, 1, 4, 5, 6, 6, 6, 8, 9, 14, 25, 49}
	bubbleSort3(c)
	fmt.Println(c)

	for {
		select {
		case <-t1:
			fmt.Println("t1定时器")
		}

	}
}
func bubbleSort3(a []int) {
	num := len(a)
	flag := true
	var count1, count2 int
	for i := num - 1; i > 0 && flag; i-- {
		for j := 0; j < i; j++ {
			flag = false
			count1++
			if a[j] > a[j+1] {
				flag = true
				count2++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Printf("循环了%d次，移动了%d次\n", count1, count2)
}
