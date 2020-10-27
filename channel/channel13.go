package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Get() string {
	duration := rand.Intn(5) + 2
	tick := time.After(time.Duration(duration) * time.Second) //等待参数duration时间后，向返回的chan里面写入当前时间。
	select {
	case <-tick:
		return fmt.Sprintf("get page %d", duration)
	}

}

func main() {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println(Get())
		wg.Done()
	}()
	go func() {
		fmt.Println(Get())
		wg.Done()
	}()
	go func() {
		fmt.Println(Get())
		wg.Done()
	}()
	wg.Wait()
}
