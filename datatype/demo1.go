package main

import (
	"fmt"
	"reflect"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(len(medals))
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
	var num uint = 0
	i := num - 1
	fmt.Println(i, reflect.TypeOf(i))
}
