package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	/*s := []int{5, 3, 3, 7, 5, 2, 9, 2}
	m := make(map[int]interface{}, 0)
	s1 := make([]int,0,len(s))
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			s1 = append(s1, v)
		}
	}
	fmt.Println(s1)*/

	defer func() {
		fmt.Printf("goroutiue num %d\n", runtime.NumGoroutine())
	}()
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	//var wg sync.WaitGroup
	//wg.Add(3)
	go func() {
		//defer func() {
		//	wg.Done()
		//}()
		duration := rand.Intn(5) + 2
		tick := time.After(time.Duration(duration) * time.Second)
		select {
		case <-tick:
			fmt.Printf("go func1 run %d\n", duration)
		case <-ctx.Done():
			fmt.Println("go func1 exit")
			return
		}
	}()

	go func() {
		//defer func() {
		//	wg.Done()
		//}()
		duration := rand.Intn(5) + 2
		tick := time.After(time.Duration(duration) * time.Second)
		select {
		case <-tick:
			fmt.Printf("go func2 run %d\n", duration)
		case <-ctx.Done():
			fmt.Println("go func2 exit")
			return
		}
	}()

	go func() {
		//defer func() {
		//	wg.Done()
		//}()

		duration := rand.Intn(5) + 2
		tick := time.After(time.Duration(duration) * time.Second)
		select {
		case <-tick:
			fmt.Printf("go func3 run %d\n", duration)
		case <-ctx.Done():
			fmt.Println("go func3 exit")
			return
		}
	}()

	//wg.Wait()

	time.Sleep(time.Second * 6)

	fmt.Println("end")
}
