package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 5, 2, 9, 1, 0}
	sort.Sort(StringSlice(a))
	sort.Ints(a)
	for _, k := range a {
		fmt.Printf("%d\n", k)
	}
	for v, _ := range a {
		fmt.Printf("%d\n", v)
	}
}
