package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

//Two ways of merging N channels in Go
func main() {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	//当channel被发送数据的协程关闭时，range就会结束，接着退出for循环
	for v := range mergeReflect(a, b, c) {
		fmt.Println(v)
	}
}

//以下是错误的，存在循环迭代变量问题：只需将变量作为参数传递，这样就可以将其复制，并且每个goroutine最终都会拥有一个与自身不同的变量。 如 merge2
func merge1(cs ...<-chan int) <-chan int {
	out := make(chan int)

	for _, c := range cs {
		go func() {
			for v := range c {
				out <- v
			}
		}()
	}
	return out
}

//错误 defer  close(out) 因为for range cs 中使用了goroutine，不能直接关闭out，关闭了就不能发送数据给out
func merge2(cs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, c := range cs {
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}
			}(c)
		}
	}()

	return out
}

func merge3(cs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(cs))
		for _, c := range cs {
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out) //channel 没有close的话，如果channel没数据的话，接受只会阻塞
	}()

	return out
}

//way one
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//way two
func mergeReflect(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(int)
		}
	}()
	return out

}

// way three
func mergeRec(chans ...<-chan int) <-chan int {
	switch len(chans) {
	case 0:
		c := make(chan int)
		close(c)
		return c
	case 1:
		return chans[0]
	default:
		m := len(chans) / 2
		return mergeTwo(
			mergeRec(chans[:m]...),
			mergeRec(chans[m:]...))
	}
}

func mergeTwo(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
