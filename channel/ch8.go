package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//把 close(c) 注释掉，程序会一直阻塞在 if data, ok := <-c; ok 那一行
		close(c)
	}()
	for {
		//ok为true说明channel没有关闭，为false说明管道已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Finished")
}
