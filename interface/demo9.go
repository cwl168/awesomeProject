package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}
type Phpher struct {
	language string
}

func (p *Phpher) debug() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p Phpher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	var c coder = &Gopher{"Go"}
	c.code()
	c.debug()

	//T类型的值不拥有所有*T指针的方法 那这样它就可能只实现更少的接口。
	var d coder = &Phpher{"Php"}
	d.code()
	d.debug()
}

//cannot use Gopher literal (type Gopher) as type coder in assignment:
//	Gopher does not implement coder (debug method has pointer receiver)
