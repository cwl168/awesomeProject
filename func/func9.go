package main

import "fmt"

func main() {
	a := func() func() {
		var i int = 10
		var j int = 10
		return func() {
			fmt.Println("i,j:%d,%d\n", i, j)
		}
	}()
	a()
}
