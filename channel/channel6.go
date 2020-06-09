package main
import (
	"fmt"
	"time"
)
/**
双向信道
 */
func main() {
	pipline := make(chan int)
	go func() {
		fmt.Println("准备发送数据: 100")
		pipline <- 100
	}()
	go func() {
		num := <-pipline
		fmt.Printf("接收到的数据是: %d", num)
	}()
	time.Sleep(100000) //单位 1ns (纳秒)
	// 主函数sleep，使得上面两个goroutine有机会执行
	//time.Sleep(time.Duration(2)*time.Second)  //休眠2秒功能
	//time.Sleep(time.Second)  //休眠1秒
}
func timestampDemo()  {
	now := time.Now()
	timestamp := now.Unix()
	fmt.Println(timestamp)  //1556101693
}
