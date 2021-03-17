package leak

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func Test3() {
	c := make(chan struct{})
	count := 0
	//time.After()表示time.Duration长的时候后返回一条time.Time类型的通道消息。
	timeout := time.After(time.Second * 6)
	done1 := make(chan struct{})

	go func() {
		for {
			select {
			case c <- struct{}{}: //往通道写值
				time.Sleep(time.Millisecond * 100)
			case <-done1:
				log.Println("退出 子goroutine...")
				return
			}
		}

	}()
	for {
		select {
		case <-c:
			count++
			fmt.Printf("count:%d\n", count)
		case <-timeout:
			fmt.Println("超时")
			return
		default:
			time.Sleep(time.Millisecond * 50)
			if count > 3 {
				done1 <- struct{}{}
				log.Println("退出 主函数...")
				return
			}
			fmt.Println("default")
		}
	}
}

func main1() {
	go func() {
		for {
			log.Println("runtime.NumGoroutine : ", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 500)
		}
	}()
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 10; i++ {
		go Test3()
	}
	time.Sleep(time.Second * 3)
}

//sayHello 是个死循环 常驻内存会引起泄露
func sayHello() {
	for {
		fmt.Println("Hello gorotine")
		time.Sleep(time.Second)
	}
}
func Leak1() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	go sayHello()
	fmt.Println("Hello main")
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//发送不接收 发送者就会被阻塞 （发送到一个没有接收者的 channel）
func Leak2() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	// Set up the pipeline.
	out := gen(2, 3)

	for n := range out {
		fmt.Println(n)              // 2
		time.Sleep(5 * time.Second) // done thing, 可能异常中断接收
		if true {                   // if err != nil
			break
		}
	}
}

func gen2(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			// select 实现 2 个 channel 的同时处理。当异常发生时，将进入 <-done 分支，实现 goroutine 退出
			select {
			case out <- n:
			case <-done:
				return

			}
		}

	}()
	return out
}

//发送不接收 发送者就会被阻塞 当接收者停止工作，发送者并不知道，还在傻傻地向下游发送数据。故而，我们需要一种机制去通知发送者。
func Leak3() {
	defer func() {
		time.Sleep(5 * time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	done := make(chan struct{})
	defer close(done)
	// Set up the pipeline.
	out := gen2(done, 2, 3)

	for n := range out {
		fmt.Println(n)              // 2
		time.Sleep(2 * time.Second) // done thing, 可能异常中断接收
		if true {                   // if err != nil
			break
		}
	}
}

//接收不发送（从没有发送者的 channel 中接收数据）
func Leak4() {
	defer func() {
		time.Sleep(5 * time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		//接收
		<-ch
	}()
}

//nil channel  向 nil channel 发送和接收数据都将会导致阻塞

func Leak5() {
	var ch chan int
	if true {
		ch = make(chan int, 1)
		ch <- 1
	}
	go func(ch chan int) {
		<-ch
	}(ch)

	defer func() {
		time.Sleep(5 * time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
}
