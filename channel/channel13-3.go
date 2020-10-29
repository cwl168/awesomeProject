package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Get 的参数由 channel 换成了 context
func Get(ctx context.Context) string {
	duration := rand.Intn(5) + 2
	tick := time.After(time.Duration(duration) * time.Second)
	select {
	case <-tick:
		return fmt.Sprintf("get page %d", duration)
	// 监控 context
	case <-ctx.Done():
		return fmt.Sprintf("cancel %d, reason:%v", duration, ctx.Err())
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	// 创建一个 context 对象
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func() {
		os.Stdin.Read(make([]byte, 1))
		// 取消 context
		cancel()
	}()
	wg.Add(3)
	go func() {
		fmt.Println(Get(ctx))
		wg.Done()
	}()
	go func() {
		fmt.Println(Get(ctx))
		wg.Done()
	}()
	go func() {
		fmt.Println(Get(ctx))
		wg.Done()
	}()
	wg.Wait()
}
