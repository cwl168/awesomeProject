package main

import (
	"fmt"
)

type Test interface{}

type Test1 interface {
	TestFunc()
}

type Structure struct {
	a int
}

func (s *Structure) TestFunc() {
	fmt.Println("Ok, Let's rock and roll!")
}

func fTest(t Test) {
	fmt.Println(t == nil)
}

func fTest1(t1 Test1) {
	fmt.Println(t1 == nil)
}

func fStructure(s *Structure) {
	fmt.Println(s == nil)
}

func main() {
	var s *Structure = nil
	fTest(s)      // false
	fTest1(s)     // false
	fStructure(s) // true
	s.TestFunc()  // Ok, Let's rock and roll!
}
