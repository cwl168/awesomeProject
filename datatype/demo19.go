package main

import (
	"fmt"
)

func main() {
	FirstNames := []string{"aaa", "bbb", "ccc"}
	LastNames := []string{"111", "222", "333"}

Loop:
	for _, firstName := range FirstNames {
		for _, lastName := range LastNames {
			fmt.Printf("Name: %s %s\n", firstName, lastName)

			if firstName == "bbb" && lastName == "111" {
				//continue Loop  //终止本次循环，继续外层下个循环，循环继续
				break Loop //跳出2层循环，循环终止
			}
		}
	}
	println("Over.")
}
