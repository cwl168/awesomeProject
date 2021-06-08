package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	var eg errgroup.Group
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			if i > 90 {
				fmt.Println("Error:", i)
				return fmt.Errorf("Error occurred: %d", i)
			}
			fmt.Println("End:", i)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
