package ConcurrentCompetition

import (
	"fmt"
	"sync"
)

func race_detector() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func race_detector2() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func race_detector3() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}

var mu sync.Mutex

func addCount() {
	m := make(map[string]string)
	var n map[int]int

	if n == nil {
		fmt.Println("nil map", len(n), len(m))
	} else {
		fmt.Println("no nil map")
	}

	m1 := make([]int, 0)
	var n1 []int

	if n1 == nil {
		fmt.Println("nil slice", len(n1), len(m1))
	} else {
		fmt.Println("no nil slice")
	}

	var wg sync.WaitGroup
	var count int
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go addGroutine(&count, &wg)
	}
	wg.Wait()
	fmt.Printf("Count:%d\n", count)
}
func addGroutine(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 2; i++ {
		mu.Lock()
		*count = *count + 1
		mu.Unlock()
	}
	wg.Done()

}
func UpdateMap() {
	c := make(chan bool)
	m := make(map[string]string, 2)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
Read at 0x00c0000aa0c0 by goroutine 9:
  command-line-arguments.addGroutine()
      /Users/caoweilin/go/src/learnDemo/ConcurrentCompetition/go12.go:20 +0x47

Previous write at 0x00c0000aa0c0 by goroutine 8:
  command-line-arguments.addGroutine()
      /Users/caoweilin/go/src/learnDemo/ConcurrentCompetition/go12.go:20 +0x5d

*/
type Class struct {
	Students sync.Map
}

func handler(key, value interface{}) bool {
	fmt.Printf("Name :%s %s\n", key, value)
	return true
}

func sysncMap() {

	class := &Class{}

	//存储值
	class.Students.Store("Zhao", "class 1")
	class.Students.Store("Qian", "class 2")
	class.Students.Store("Sun", "class 3")

	//遍历，传入一个函数，遍历的时候函数返回false则停止遍历
	class.Students.Range(handler)

	//查询
	if _, ok := class.Students.Load("Li"); !ok {
		fmt.Println("-->Li not found")
	}

	//查询或者追加
	_, loaded := class.Students.LoadOrStore("Li", "class 4")
	if loaded {
		fmt.Println("-->Load Li success")
	} else {
		fmt.Println("-->Store Li success")
	}

	//删除
	class.Students.Delete("Sun")
	fmt.Println("-->Delete Sun success")

	//遍历
	class.Students.Range(handler)

}
