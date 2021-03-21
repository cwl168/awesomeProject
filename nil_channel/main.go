package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

outer:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("interruppted after 1s tick")
		case <-ctx.Done():
			fmt.Println("We're done here!")
			break outer
		}
	}
}
