package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//https://www.godesignpatterns.com/2014/05/nil-channels-always-block.html
//https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308

/**
The code above will cause a very subtle bug in your program if one of the two channels is closed. Since a closed channel will never block,
if ch1 is closed, ch1 will always be ready to receive on subsequent iterations through the loop.
As a result, the program will enter an infinite spin loop, and case <-ch2: may or may not be given a chance to run later on.
*/

/**
Given a nil channel c:
<-c receiving from c blocks forever
c <- v sending into c blocks forever
close(c) closing c panics
*/
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
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
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					log.Println("a is done")
					adone = true
					continue
				}
				log.Printf("add from a %d\n", v)
				c <- v
			case v, ok := <-b:
				if !ok {
					log.Println("b is done")
					bdone = true
					continue
				}
				log.Printf("add from b %d\n", v)
				c <- v
			}
		}
	}()
	return c
}
func merge2(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					log.Println("a is done")
					a = nil
					continue
				}
				log.Printf("add from a %d\n", v)
				c <- v
			case v, ok := <-b:
				if !ok {
					log.Println("b is done")
					b = nil
					continue
				}
				log.Printf("add from b %d\n", v)
				c <- v
			}
		}
	}()
	return c
}
