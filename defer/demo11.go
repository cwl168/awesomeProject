package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (a interface{}) {
	defer func() {
		if err := recover(); err != nil {
			a = err
		}
	}()
	panic(55)
}
