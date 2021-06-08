package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func Get(ctx context.Context) string {
	duration := rand.Intn(5) + 2
	tick := time.After(time.Duration(duration) * time.Second)
	select {
	case <-tick:
		return fmt.Sprintf("get page %d", duration)
	case <-ctx.Done():
		return fmt.Sprintf("cancel %d, reason:%v", duration, ctx.Err())
	}
}

func main() {
	defer func() {
		fmt.Printf("goroutiue num %d\n", runtime.NumGoroutine())
	}()
	rand.Seed(time.Now().Unix())
	//ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func() {
		os.Stdin.Read(make([]byte, 1))
		context.Canceled = errors.New("手动取消请求")
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
