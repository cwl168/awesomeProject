package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var key string = "name"

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), 3000*time.Millisecond) //WithTimeout设置超时时间
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(500*time.Millisecond)) //WithDeadline是设置具体时间点

	wg.Add(3)
	go func(ctx context.Context) {
		defer func() {
			wg.Done()
		}()
		valueCtx := context.WithValue(ctx, key, "request1")
		Get1(valueCtx)
	}(ctx)

	go func(ctx context.Context) {
		defer func() {
			wg.Done()
		}()
		valueCtx := context.WithValue(ctx, key, "request2")
		Get1(valueCtx)
	}(ctx)

	go func(ctx context.Context) {
		defer func() {
			wg.Done()
		}()
		valueCtx := context.WithValue(ctx, key, "request3")
		Get1(valueCtx)
	}(ctx)

	wg.Wait()
	fmt.Println("close")
	fmt.Println("end")
}
func Get1(ctx context.Context) {
	duration := rand.Intn(5) + 2
	tick := time.After(time.Duration(duration) * time.Second) //表示time.Duration长的时候后返回一条time.Time类型的通道消息
	select {
	case <-tick:
		fmt.Printf("Get %s %d\n", ctx.Value(key), duration)
	case <-ctx.Done():
		fmt.Printf("%s 请求关闭...  %s\n", ctx.Value(key), ctx.Err())
		return

	}
}
