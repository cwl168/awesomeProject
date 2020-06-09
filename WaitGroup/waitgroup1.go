package main
/*
Add：初始值为0，你传入的值会往计数器上加，这里直接传入你子协程的数量
Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
Wait：阻塞当前协程，直到实例里的计数器归零。
 */
import (
	"fmt"
	"sync"
)

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
}
