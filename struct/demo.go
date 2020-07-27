package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//1、顺序初始化，必须每个成员都初始化
	//var s1 Student = Student{1, "mike", 'm', 18, "sz"}
	//s2 := Student{2, "yoyo", 'f', 20, "sz"}
	//s3 := Student{2, "tom", 'm', 20} //err, too few values in struct initializer

	//2、指定初始化某个成员，没有初始化的成员为零值
	s4 := Student{id: 2, name: "lily"}
	fmt.Println(s4)
}
