package main

import (
	"bytes"
	"io"
)

func main() {
	var buf *bytes.Buffer
	test(buf)
}
func test(w io.Writer) {
	if w != nil {
		w.Write([]byte("ok"))
	}
}
