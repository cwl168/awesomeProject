package main

import (
	"fmt"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	f := 3.14
	fmt.Printf("f type = %T\n", f)
	var a interface{}
	var u = User{
		ID:   1,
		Name: "curry",
	}
	a = u.ID
	fmt.Printf("%T\t%#[1]v\n", a)
	fmt.Printf("%s\n", a)
}
