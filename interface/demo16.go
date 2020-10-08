package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf *bytes.Buffer
	fmt.Printf("%T\t%[1]v\n", buf)
	out := io.Writer(buf)
	fmt.Printf("%T\t%[1]v\n", out)
	fmt.Println(out == nil)

	/*var buf io.Writer
	fmt.Printf("%T\t%[1]v\n", buf)
	out := io.Writer(buf)
	fmt.Printf("%T\t%[1]v\n", out)
	fmt.Println(out == nil)*/
}
