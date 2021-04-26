package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
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
	rand.Seed(time.Now().Unix())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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
