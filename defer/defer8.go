package main


func test() (x int) {
	defer println("defer")
	return 200
}

func main() {
	println(test())
}
/**
defer 是return 后才调用的
 */