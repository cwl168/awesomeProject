package main

import (
	"fmt"
)

func hello(i *int) *int {
	defer func() {
		*i = 19
	}()
	return i
}

func main() {
	i := 10
	j := hello(&i)
	fmt.Println(i, *j)
}

func hello2(i *int) (j int) {
	defer func() {
		j = 19
	}()
	j = *i
	return j
}

func hello3(i *int) (j int) {
	defer func() {
		j = 19
	}()
	return *i
}
