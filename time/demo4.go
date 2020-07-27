package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	w.Close()



	dir, err := os.Getwd()
	if err != nil {
		 fmt.Errorf("os.Getwd err: %v", err)
	}
	fmt.Println(dir)
	timer1 := time.NewTimer(time.Second * 5)
	t1 := time.Now() //当前时间
	fmt.Printf("t1: %v\n", t1)

	for {
		select {
		case t2:=<-timer1.C:
			fmt.Printf("t2: %v\n", t2)
			timer1.Reset(time.Second * 5)
		}
	}

}
