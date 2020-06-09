package main

import (
	"fmt"
)

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]//slice[start:index:max],它的容量就是(max - start)
	fmt.Printf("myslice为 %d, 其长度为: %d, 其容量为: %d\n\n", myslice, len(myslice),cap(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Println("myslice =", myslice)
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}
