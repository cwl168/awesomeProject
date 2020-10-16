package main

import (
	"fmt"
	"os"
)

func main() {
	worklist := make(chan []string)
	worklist <- os.Args[1:]

	for list := range worklist {
		fmt.Println(list)
	}
}
