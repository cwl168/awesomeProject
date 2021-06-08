package main

import (
	"fmt"
)

var m = make(map[string]interface{})

//将slice转化为字符串作为map的key
func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int { return m[k(list)] }

func main() {

	Add([]string{"1", "2", "3"})
	Add([]string{"1", "2", "3"})
	Add([]string{"1", "2", "3", "4"})
	Add([]string{"1", "2", "3"})

	if k([]string{"1", "2", "3"}) == k([]string{"1", "2", "3"}) {
		fmt.Println("equal")
	}

	fmt.Println(m)
	fmt.Println(Count([]string{"1", "2", "3"}))

	fmt.Println(len(m))
	for k, v := range m {
		fmt.Println(k, v)
	}

}

// <===================OUTPUT===================>
// 3
// ["1" "2" "3" "4"] 1
// ["1" "2" "3"] 3
