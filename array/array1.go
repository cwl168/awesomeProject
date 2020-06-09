package main
import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Printf("%d\n",arr)
	var arr00 [3]int = [3]int{1,2,3}
	fmt.Printf("%d\n",arr00)
	arr01 := [...]int{1, 2, 3}
	arr02 := [...]int{1, 2, 3, 4}
	fmt.Printf("%d 的类型是: %T\n", arr01, arr01)
	fmt.Printf("%d 的类型是: %T", arr02, arr02)
}
