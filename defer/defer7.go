package main

import "log"

func main() {
	defer dd()()
	log.Println("789")
}

func dd() func(){
	log.Println("123")
	return func() {
		log.Println("456")
	}
}
/**
defter  函数1函数2
执行1
--------
函数1->执行1->函数2
defer只延迟最后一个函数

2020/05/29 18:21:41 123
2020/05/29 18:21:41 789
2020/05/29 18:21:41 456

*/