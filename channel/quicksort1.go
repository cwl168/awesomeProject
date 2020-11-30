package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

//左边比哨兵小，右边比哨兵大，并返回最终哨兵的位置
func partition(a []int, low, high int) int {
	item := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < item {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	//交换 处理哨兵
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func quickSort(a []int, low, high int) {
	if low >= high {
		return
	}

	p := partition(a, low, high)
	quickSort(a, low, p-1)
	quickSort(a, p+1, high)
}

func goQuickSort(a []int, low, high int, cdone chan struct{}, depth int) {
	if low >= high {
		cdone <- struct{}{}
		return
	}
	//控制goroutine数量
	depth--
	p := partition(a, low, high)
	if depth > 0 {
		cdone := make(chan struct{}, 2)
		go goQuickSort(a, low, p-1, cdone, depth)
		go goQuickSort(a, p+1, high, cdone, depth)

		<-cdone
		<-cdone
	} else {
		quickSort(a, low, p-1)
		quickSort(a, p+1, high)
	}
	//排序完成 通知主groutinue
	mdone <- struct{}{}
}

const num = 1000000

var mdone = make(chan struct{})

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	rand.Seed(time.Now().UnixNano())
	slice1, slice2 := make([]int, 0, num), make([]int, 0, num)
	/*for i := 0; i < num; i++ {
		val := rand.Intn(1000)
		slice1 = append(slice1, val)
		slice2 = append(slice2, val)
	}*/
	slice1 = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	slice2 = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	start := time.Now()
	quickSort(slice1, 0, len(slice1)-1)
	fmt.Println("非并发版: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(slice1) {
		fmt.Println("非有序")
	}
	fmt.Println(slice1)

	start = time.Now()
	go goQuickSort(slice2, 0, len(slice2)-1, mdone, 1)
	<-mdone
	fmt.Println("并发版: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(slice2) {
		fmt.Println("非有序")
	}
}
