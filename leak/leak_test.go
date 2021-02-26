package main

import (
	"github.com/fortytw2/leaktest"
	"testing"
	"time"
)

//go test -v
//go test -run TestTest3 -v

// 自动化测试leak  检测goroutine内存泄露
func TestTest3(t *testing.T) {
	defer leaktest.Check(t)()
	Test3()
}

func TestPool(t *testing.T) {
	defer leaktest.Check(t)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}
