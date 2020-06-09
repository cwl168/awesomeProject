package main

import "fmt"

func main() {
	/*ch := make(chan int,1)
	for i := 0; i < 10; i++{
		select {
		case x := <- ch:
			fmt.Println(x)
		case ch <- i:

		}
	}*/
	/*timeout := make(chan bool, 1)
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	default:
		fmt.Println("default case is running")
	}*/
	//超时
	/*timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3e9) // sleep 3 seconds
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	}*/
	ch := make(chan int, 2)
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}
