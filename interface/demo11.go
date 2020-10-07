package main

import (
	"bytes"
	//"bytes"
	"fmt"
	"io"
)

func main() {
	//var buf io.Writer
	var buf *bytes.Buffer
	fmt.Printf("main %T,%v\n", buf, buf)
	test(buf)
}

//在main()中，我们首先声明了一个buf变量，类型是*bytes.Buffer指针类型，
//在调用函数test()的时候，参数w会被赋值为动态类型为*bytes.Buffer，动态值为nil，也就是w是一个包含了空指针值的非空接口。
//那么在w != nil判断时，这个等式便是成立的，不过这里也从侧面反映出一个现象，就是这种传参都是值拷贝
func test(w io.Writer) {
	fmt.Printf("test %T,%v\n", w, w)
	//w是一个包含了空指针值的非空接口
	if w != nil {
		w.Write([]byte("ok"))
	} else {
		fmt.Println("空接口！")
	}
}
