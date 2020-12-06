package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}
type Student struct {
	name string
	age  int
}

var P *Person
var S *Student
var list = make(map[string]interface{})

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func (p Student) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func addMap(k string, v interface{}) {
	if _, ok := list[k]; !ok {
		list[k] = v
	}
}

func main() {
	P := &Person{"Dennis", 70}
	S := &Student{"Cwllin", 60}
	addMap("p", &P)
	addMap("S", &S)

	for index, element := range list {
		switch value := element.(type) {
		case **Student:
			fmt.Printf("list[%s] is a Student and its value is [%v]\n", index, *value)
			break
		case **Person:
			fmt.Printf("list[%s] is a Person and its value is [%s]\n", index, *value)
			break
		default:
			fmt.Printf("list[%s] is of a different type", index)
		}
	}
}
