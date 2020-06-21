package main

import "fmt"

//定义一个闭包
func DoWork(name string) func(int) int {
	//设置一个保存进度的状态
	var prop int = 0

	f := func(hours int) int {
		prop += hours
		fmt.Printf("%s工作了%d小时\n", name, hours)
		return prop
	}

	return f

}

func main(){
	//使用闭包
	f1 := DoWork("fun")
	f2 := DoWork("joker")

	p1 := f1(22)
	p2 := f2(8)

	p1 = f1(7)
	p2 = f2(12)

	fmt.Println("\n")
	fmt.Printf("fun总共工作了%d小时\n", p1)
	fmt.Printf("joker总共工作了%d小时\n", p2)

}
