package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("~/go/src/learnDemo/file/1.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if(err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("body type = %T\n", f)
}
