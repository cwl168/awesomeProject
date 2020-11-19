package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/**
 * 作用：并发排序
 * 参数：
 *   arr 待排序数组
 *   channel： 结果管道
 *   level：   当前层级
 *	 threads： 线程总数
 * 返回值： 无
 */
func QuickSort(arr []int, channel chan int, level int, threads int) {
	level = level * 2 //每加深一个级，多*2个线程
	length := len(arr)

	if length <= 0 {
		close(channel)
		return
	} else if length <= 1 {
		channel <- arr[0]
		close(channel)
		return
	} else {
		low := make([]int, 0)
		mid := make([]int, 0)
		high := make([]int, 0)

		rand.Seed(time.Now().Unix())
		base := arr[rand.Intn(length)]

		for i := 0; i < length; i++ {
			if arr[i] < base {
				low = append(low, arr[i])
			} else if arr[i] == base {
				mid = append(mid, arr[i])
			} else {
				high = append(high, arr[i])
			}
		}

		left_chan := make(chan int, len(low)) // 注意这里必修加buffer。 否则会死锁。。为什么？？？？
		right_chan := make(chan int, len(high))

		if level <= threads { //如果线程超过执行数量，顺序调用，否则并发调用
			go QuickSort(low, left_chan, level, threads)
			go QuickSort(high, right_chan, level, threads)
		} else {
			QuickSort(low, left_chan, level, threads)
			QuickSort(high, right_chan, level, threads)
		}

		for v := range left_chan {
			channel <- v
		}
		for _, v := range mid {
			channel <- v
		}
		for v := range right_chan {
			channel <- v
		}

		close(channel)
		return
	}
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	data := []int{3, 6, 23, 7, 2, 4, 9, 13}
	channel := make(chan int)
	go QuickSort(data, channel, 1, 10) // 这里必须是go，否则会一直阻塞，导致死锁
	for v := range channel {
		fmt.Println(v)
	}
}
