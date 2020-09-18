package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	//w.Circle.X = 8
	w.X = 8
	fmt.Println(w)
}
