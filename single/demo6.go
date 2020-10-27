package main

import (
	"fmt"
	"sync"
)

type example3 struct {
	name string
}

var instance3 *example3
var once sync.Once

func GetInstance3() *example3 {

	once.Do(func() {
		instance3 = new(example3)
		instance3.name = "第一次赋值单例"
	})
	return instance3
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			obj := GetInstance3()
			fmt.Printf("%d,%p\n", i, obj)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
