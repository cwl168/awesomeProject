package main

import "fmt"

func main() {
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(len(medals))
	var num uint = 0
	for i := num; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}
