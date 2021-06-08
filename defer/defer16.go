package main

import (
	"log"
	"time"
)

//在Go语言中，是存在一些无法恢复的致命错误方法的，如fatalthrow方法和fatalpanic方法等，它们一般在并发写入map等处理时抛出，需要谨慎。(这种错误无法捕获)
func main() {
	m := make(map[int]string)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("recover: %v", e)
				}
			}()
			m[i] = "Go编程之旅：一起用Go做项目"
		}()
	}
	time.Sleep(5 * time.Second)
}
