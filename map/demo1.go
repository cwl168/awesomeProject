package main

import "fmt"


func main() {
	var m1  = map[int]interface{}{1:2, 2:4}
	var m2  = map[int]interface{}{1:"hello", 2:"world"}
	fmt.Println(m2)
	fmt.Println(m1)

	fmt.Println(m1)
	fmt.Println(m2)

	//make 与 new
	p := new([2]int) //p == nil; with len and cap 0
	(*p)[0] = 18  //because p is a nil pointer, with len and cap 0
	(*p)[1] = 19

	fmt.Println(p)
	fmt.Printf("p type %T\n",p)
	p1 := new(int)
	*p1 = 2
	fmt.Printf("p1 type %T\n",p1)
	*p1 = '5'
	fmt.Printf("p1 type %T  %[1]v\n",p1)

	v := make([]int, 10, 50) // v is initialed with len 10, cap 50
	fmt.Println(v)
	fmt.Printf("v  type %T \n",v)



	q := make(map[int]interface{},2) // v is initialed with len 10, cap 50
	fmt.Println(q)
	fmt.Printf("q type %T\n",q)
}
