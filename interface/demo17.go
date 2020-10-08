package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func main() {
	var buf *bytes.Buffer
	//var buf io.Writer
	buf = new(bytes.Buffer)
	f(buf)
}

func IsValidObject(value interface{}) bool {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val.IsValid()
}

func f(out io.Writer) {
	fmt.Printf("%T\t%[1]v\n", out)
	if !IsValidObject(out) {
		return
	}
	// ...do something...
	//if out != nil {
	fmt.Println("out write...")
	out.Write([]byte("done!\n"))
	//}
}
