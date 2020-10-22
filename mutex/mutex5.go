// demo01.go
package main

import (
	"fmt"
	"sync"
)

var a int
var mu sync.Mutex

/*func run() {
	a++
}*/
func run() {
	mu.Lock()
	defer mu.Unlock()
	a++
}

func main() {
	var wg sync.WaitGroup
	num := 100
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			run()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("a = %d\n", a)
}
