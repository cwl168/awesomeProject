package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	//var buf os.File
	//buf.Write([]byte("done!\n"))
	//var buf io.Writer
	var buf *bytes.Buffer
	fmt.Printf("f %T,%v\n", buf, buf)
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	//out的动态值是nil。，它的动态类型是*bytes.Buffer，out变量是一个包含空指针值的非空接口,out!=nil的结果依然是true
	// ...do something...
	fmt.Printf("f %T,%v\n", out, out)
	if out != nil {
		fmt.Println("out write...")
		out.Write([]byte("done!\n"))
	}
}
