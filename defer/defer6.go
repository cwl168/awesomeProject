package main

func test1() (x int)  {
	defer func() {
		x = 100
	}()

	return 200
}

func main() {
	println(test1())
}
