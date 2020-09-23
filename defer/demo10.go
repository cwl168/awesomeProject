package main

import (
	"fmt"
	"os"
)

func deferInLoops() {
	for i := 0; i < 10; i++ {
		f, _ := os.Open("/etc/hosts")
		fmt.Println(&f)
		defer f.Close()
	}
}
func main() {
	deferInLoops()
}
