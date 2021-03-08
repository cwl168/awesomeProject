package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	chanLen := 10
	goroutineNum := 4
	var wg sync.WaitGroup

	//定义buffer通道
	ch := make(chan string, chanLen)

	wg.Add(goroutineNum)

	for i := 1; i <= chanLen; i++ {
		ch <- fmt.Sprintf("task %d", i)
	}

	//4个goroutine 消费
	for i := 1; i <= goroutineNum; i++ {
		go worker(ch, i, &wg)
	}

	close(ch)

	wg.Wait()

}

func worker(ch chan string, worker int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Printf("Worker:%d,Done\n", worker)

	}()
	/*for {
		task, ok := <-ch
		if !ok {
			fmt.Printf("Worker:%d,Shutting Down\n", worker)
			return
		}

		fmt.Println(task + " start working")

		time.Sleep(2 * time.Second)

		fmt.Printf("Worker:%d,Completed %s\n", worker, task)

	}*/

	for task := range ch {
		fmt.Println(task + " start working")
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker:%d,Completed %s\n", worker, task)
	}
	fmt.Printf("Worker:%d,Shutting Down\n", worker)
}
