package main

import (
	"fmt"
	"image"
)

type st struct {
	name string
	age  int
}
type info struct {
	result int
}

func main() {
	icons := make(map[string]image.Image)
	if icons != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}

	//var m map[int]string
	m := make(map[int]string)
	m[1] = "123"
	fmt.Println(m)

	var s []int
	s = append(s, 1)
	fmt.Println(s)

	var data info
	var err error
	data.result, err = work() //不能使用简短声明来设置字段的值
	fmt.Println(data, err)

}
func work() (int, error) {
	return 3, nil
}
