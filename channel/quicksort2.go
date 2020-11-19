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
	arr := []int{7, 3, 9, 4, 6}
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
	pivotIndex := 0
	pivot := values[pivotIndex]
	storeIndex := pivotIndex + 1
	for i := pivotIndex + 1; i < len(values); i++ {
		if values[i] < pivot {
			if i != storeIndex {
				values[i], values[storeIndex] = values[storeIndex], values[i]
			}
			storeIndex++
		}
	}
	values[pivotIndex], values[storeIndex-1] = values[storeIndex-1], values[pivotIndex]

	wg.Add(2)
	go quickSort(values[0:storeIndex-1], wg)
	go quickSort(values[storeIndex:], wg)
	wg.Done()
}
