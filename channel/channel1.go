package main

import (
	"fmt"
	"os"
)

//var ch9 = make(chan int)
var ch9 = make(chan map[int]int)

func mm2() {
	for i := 1; i < 10; i++ {
		result := make(map[int]int)
		result[i] = 8 * i
		ch9 <- result
	}
}
func main() {
	in, err := os.Open("infile")
	out, err := os.Create("outfile")
	fmt.Println(in)
	fmt.Println(out)
	fmt.Println(err)
	ch := make(chan int)
	<-ch
	//go mm2()
	//for v:=range ch9{
	//	fmt.Println(v)
	//}
	////如果通道未关闭，则报错
	//for {
	//	if data, ok := <-ch9; ok {
	//		for k,v := range data {
	//			fmt.Printf("k->%d ,v->%d\n",k,v)
	//		}
	//	} else {
	//		break
	//	}
	//}
	//如果通道未关闭，则报错
	/*for v:=range ch9{
		fmt.Println(v)
	}*/
	//读的次数大于写，如果通道关闭，则默认值，如果未关闭，则报错
	/*for i:=0;i<100;i++{
		fmt.Println(<-ch9)
	}*/

}
