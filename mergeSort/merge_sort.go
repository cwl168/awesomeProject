package mergeSort

import (
	"fmt"
	"sort"
)

//同步版
func mergeSort1(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	fmt.Printf("length:%d,num:%d\n", length, num)
	left := mergeSort1(r[:num])
	right := mergeSort1(r[num:])
	return merge(left, right)
}

//做合并的left right 都是有序的
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	//剩余元素追加
	fmt.Printf("l:%d,len(left):%d,r:%d,len(right):%d,%v\n", l, len(left), r, len(right), result)
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	fmt.Println(l, left[l:], r, right[r:], result)
	fmt.Println()

	return
}

//并发版
func mergeSort2() {
	//生成数据
	a := produceNum(1, 4, 8, 2, 19, 5)
	b := produceNum(0, 29, 43, 1, 7, 9)
	c := produceNum(45, 15, 20, 11, 25, 13)
	d := produceNum(60, 75, 55, 12, 31, 35)
	f := produceNum(70, 52, 65, 80, 59, 37)

	//排序
	a_sort := sortNum(a)
	b_sort := sortNum(b)
	c_sort := sortNum(c)
	d_sort := sortNum(d)
	f_sort := sortNum(f)

	//合并
	ch := mergeSort(a_sort, b_sort, c_sort, d_sort, f_sort)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("end")
}

//两两合并
func mergeSort(chanSlice ...<-chan int) <-chan int {
	length := len(chanSlice)
	if len(chanSlice) == 1 {
		return chanSlice[0]
	}
	num := length / 2
	left := mergeSort(chanSlice[:num]...)
	right := mergeSort(chanSlice[num:]...)
	return mergeSortTwo(left, right)
}
func mergeSortTwo(left, right <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-left
		v2, ok2 := <-right
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-left
			} else {
				out <- v2
				v2, ok2 = <-right
			}
		}
		close(out)
		fmt.Println("mergeSort data end")
	}()
	return out
}
func sortNum(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		s := make([]int, 0)
		for v := range ch {
			s = append(s, v)
		}
		sort.Ints(s)
		for _, v := range s {
			out <- v
		}
		close(out)
	}()
	return out
}
func produceNum(s ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range s {
			out <- v
		}
		close(out)
	}()
	return out
}
