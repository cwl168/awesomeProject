package main

import (
	"fmt"
	"time"
)

func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")

		t := time.NewTicker(time.Millisecond * 500)

		// Using stop channel explicit exit
		for {
			select {
			case v := <-stopCh:
				fmt.Printf("%T\n", v) //返回nil 默认值
				fmt.Println("Recv stop signal")
				return
			case <-t.C:
				fmt.Println("Working .")
			}
		}
	}()
	return
}

func main() {

	stopCh := make(chan struct{})
	worker(stopCh)

	time.Sleep(time.Second * 2)
	fmt.Println("close sending")
	close(stopCh) //channel 没有close的话，如果channel没数据的话，接受只会阻塞,等待着发送者发送数据,如果数据满了，发送数据端将阻塞等待接收 ,channel close后，发送数据将会触发 panic 异常，还能接受channel中的数据，等channel中的数据为空，返回nil,

	// Wait some print
	time.Sleep(time.Second)
	fmt.Println("main exit")
}

// 协程处理1个通道，并且是读时，协程优先使用for-range，当channel被发送数据的协程关闭时，range就会结束，接着退出for循环, 该读取协程自动退出。
// x,ok可以处理多个读通道关闭，需要关闭当前使用for-select的协程。
// 多个channel都关闭才能退出,利用select的一个特性，select不会在nil的通道上进行等待
// 显式关闭通道stopCh可以处理主动通知协程退出的场景。
// context

//goroutine 接受channel数据，或是 监听channel关闭事件
