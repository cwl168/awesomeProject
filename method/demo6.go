package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add //方法表达式
	} else {
		op = Point.Sub //方法表达式
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		fmt.Println(path[i])
		path[i] = op(path[i], offset)
	}
}

//变量op代表Point类型的addition或者subtraction方法，Path.TranslateBy方法会为其Path数组中的每一个Point来调用对应的方法
func main() {
	var p Path
	p = []Point{{2, 3}, {4, 5}}
	p.TranslateBy(Point{1, 2}, true)
	fmt.Println(p)

	var sum int
	vals := []int{1, 2, 3}
	for _, val := range vals {
		sum += val
	}
	fmt.Println(sum)
}
