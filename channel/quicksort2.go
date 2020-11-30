package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"
)

const num = 100000

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	/*arr := make([]int, 0, num)
	for i := 0; i < num; i++ {
		val, _ := rand.Int(rand.Reader, big.NewInt(1000))
		arr = append(arr, int(val.Int64()))
	}*/
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go quickSort(arr, &wg)

	wg.Wait()
	fmt.Printf("%dms elapsed.\n", time.Since(start).Milliseconds())

	if !sort.IntsAreSorted(arr) {
		fmt.Errorf("非有序")
	}
}

func quickSort(values []int, wg *sync.WaitGroup) {
	if len(values) <= 1 {
		wg.Done()
		return
	}
	h := len(values) - 1
	pivotIndex := 0
	pivot := values[h]
	storeIndex := pivotIndex - 1
	for i := pivotIndex; i < len(values); i++ {
		if values[i] < pivot {
			storeIndex++
			values[i], values[storeIndex] = values[storeIndex], values[i]
		}
	}
	values[h], values[storeIndex+1] = values[storeIndex+1], values[h]

	wg.Add(2)
	go quickSort(values[0:storeIndex+1], wg)
	go quickSort(values[storeIndex+2:], wg)
	wg.Done()
}

/*func quickSort2(values []int) {
	if len(values) <= 1 {
		return
	}
	pivotIndex := 0
	pivot := values[pivotIndex]
	storeIndex := pivotIndex + 1
	for i := pivotIndex + 1; i < len(values); i++ {
		count1++
		if values[i] < pivot {
			if i != storeIndex {
				count2++
				values[i], values[storeIndex] = values[storeIndex], values[i]
			}
			storeIndex++
		}
	}
	values[pivotIndex], values[storeIndex-1] = values[storeIndex-1], values[pivotIndex]

	quickSort(values[0 : storeIndex-1])
	quickSort(values[storeIndex:])
}*/
