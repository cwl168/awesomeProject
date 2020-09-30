package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer //类型和值的部分都是nil

	if w == nil {
		//w.Write([]byte("hello")) //panic: nil pointer dereference
		fmt.Println("空")
	} else {
		fmt.Println("不空")
	}
	w = os.Stdout         //动态类型被设 为*os.Stdout指针的类型描述符，它的动态值持有os.Stdout的拷贝；
	w = new(bytes.Buffer) //动态类型是*bytes.Buffer,动态值是一个指向新分配的缓冲区的指针
	fmt.Printf("%T\n", w) // 打印动态类型
	w = nil
	fmt.Println(w)

	var buf *bytes.Buffer
	test(buf)
}
func test(w io.Writer) {
	fmt.Printf("%T\n", w)
	//w会被赋值为动态类型为*bytes.Buffer，动态值为nil，也就是w是一个包含了空指针值的非空接口。那么在w != nil判断时，这个等式便是成立的
	if w != nil {
		w.Write([]byte("ok"))
	}
}
