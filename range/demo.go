package main

import "fmt"

func main() {
	//切片
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Println(i)
		fmt.Println(v)
	}
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)
	for index, value := range slice {
		fmt.Println(&index, &value)
		myMap[index] = &value
	}
	fmt.Println("=====new map=====")
	for k, v := range myMap {
		fmt.Printf("%d => %d\n", k, *v)
	}
}
