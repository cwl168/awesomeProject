package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func Get(done <-chan int) string {
	duration := rand.Intn(5) + 2
	tick := time.After(time.Duration(duration) * time.Second) //等待参数duration时间后，向返回的chan里面写入当前时间。
	select {
	case <-tick:
		return fmt.Sprintf("get page %d", duration)
	case d := <-done:
		return fmt.Sprintf("cancel %d,value %d", duration, d)
	}
}
func main() {
	rand.Seed(time.Now().Unix())
	// 创建一个 channel
	done := make(chan int)
	var wg sync.WaitGroup
	go func() {
		os.Stdin.Read(make([]byte, 1))
		// 因为我们发起了 3 次 Get 请求，所以要发送 3 次数据到 done channel
		/*for n := 0; n < 3; n++ {
			done <- struct{}{}
		}*/
		//如果从一个已经关闭的 channel 读取数据，程序会立即返回而不阻塞。
		close(done)
	}()
	wg.Add(3)
	go func() {
		fmt.Println(Get(done))
		wg.Done()
	}()
	go func() {
		fmt.Println(Get(done))
		wg.Done()
	}()
	go func() {
		fmt.Println(Get(done))
		wg.Done()
	}()
	wg.Wait()
}
