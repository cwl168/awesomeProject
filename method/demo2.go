package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	pptr := &p

	//不论是接收器的实际参数和其接收器的形式参数相同，比如两者都是类型T或者都是类型*T：
	p.Distance(q)
	pptr.ScaleBy(2)

	//或者接收器形参是类型T，但接收器实参是类型*T，这种情况下编译器会隐式地为我们取变量的地址：
	p.ScaleBy(2)

	//或者接收器形参是类型*T，实参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量：
	pptr.Distance(q)
	fmt.Printf("main Distance &p = %p\n", pptr)
}
