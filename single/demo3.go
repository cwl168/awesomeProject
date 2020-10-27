package main

import (
	"fmt"
	"sync"
)

type example struct {
	name string
}

var instance *example

var mux sync.Mutex

func GetInstance() *example {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = &example{}
	}
	return instance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			obj := GetInstance()
			fmt.Printf("%d,%p\n", i, obj)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
