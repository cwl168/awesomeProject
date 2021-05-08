package main

import (
	"fmt"
	"runtime"
)

type People struct{ Name string }

func (p *People) String() string {
	return fmt.Sprintf("printffff: %v\n", p.Name)
}

func main() {
	var numbers = make([]int, 3, 5)

	var numbers3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3[2:]) //索引从2
	fmt.Println(numbers3[:2]) //索引从 2-1

	n4 := numbers3[1:2:2]
	n4[0] = 11
	n4 = append(n4, 12)
	fmt.Printf("n4:%v\n", n4)
	fmt.Printf("numbers3:%v\n", numbers3)

	var slice1 = numbers3[1:4] //2 3 4

	var slice2 = slice1[1:3] //3 4

	s := make([]int, 2, 4)
	s = append(s, 3)
	fmt.Println(s, len(s), cap(s), s[0], s[1])

	for i, x := range slice2 {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	printSlice(numbers)

	p := &People{Name: "cwl"}
	fmt.Println(p)

	fmt.Println(runtime.NumCPU())

	sss := []int{8, 4, 5, 7, 1, 3, 2, 6}
	Quick2Sort(sss)
	fmt.Println(sss)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}
