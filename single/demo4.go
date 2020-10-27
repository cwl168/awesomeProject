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
	if instance == nil { // 单例没被实例化，才会加锁
		mux.Lock()
		defer mux.Unlock()
		//假设现在两个线程来请求GetInstance，A、B线程同时发现singleton实例对象为空，因为我们在第一次check方法上没有加锁，然后A线程率先获得锁，进入同步代码块，new了一个singleton实例对象，之后释放锁，接着B线程获得了这个锁，发现singleton实例对象已经被创建了，就直接释放锁，退出同步代码块。防止这种情况
		if instance == nil { // 单例没被实例化才会创建
			instance = &example{}
		}
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
