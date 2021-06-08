package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 8, 7, 1, 3, 5, 6, 4}
	Quick3Sort(arr, 0, len(arr)-1)
	fmt.Println("------")
	fmt.Println(arr)

	s := "sass"
	fmt.Println(len(s), s[1])
	bs := ([]byte)(s)
	bs[1] = 'b'
	fmt.Printf("%c\n", bs[1])
}

// 第二种写法
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

// 第三种写法
func Quick3Sort(a []int, left int, right int) {
	if left >= right {
		return
	}
	//基准索引与元素
	pivotIndex := left
	pivotValue := a[left]
	for i := left + 1; i <= right; i++ {
		//找比基准元素小的
		if pivotValue >= a[i] {
			//移动到下个元素
			pivotIndex++
			//分割位定位++
			a[i], a[pivotIndex] = a[pivotIndex], a[i]
		}
	}
	//起始位和分割位
	a[left], a[pivotIndex] = a[pivotIndex], a[left]

	Quick3Sort(a, left, pivotIndex-1)
	Quick3Sort(a, pivotIndex+1, right)

}
