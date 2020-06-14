package main

/**
多个defer的执行顺序为“后进先出”；
defer、return、返回值三者的执行逻辑应该是：return设置返回值（这里分为无名返回值和有名返回值），defer执行，函数带着返回值退出；
 */

//有名返回值
func test1() (x int)  {
	defer func() {
		x = 100
	}()

	return 200
}

func main() {
	println(test1())
}
