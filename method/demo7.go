package main

import (
	"fmt"
	"learnDemo/method/person"
)

func main() {
	var p = person.NewPerson("cwl")
	(*p).Setage(31)
	fmt.Println(*p)
}
