//效率排序:
//append> strings.Join() > + > bytes.writestring > fmt
package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

var loop = 100000

func main() {
	var s string
	s1 := "hello"
	s2 := "world"
	var start time.Time
	//加号+连接
	start = time.Now()
	for i := 0; i < loop; i++ {
		s1 += s2
	}
	fmt.Println("+连接方法:", time.Since(start))
	//append连接
	s1 = "hello"
	s2 = "world"
	start = time.Now()
	for i := 0; i < loop; i++ {
		//[]byte(s1) s1转化为切片
		s = string(append([]byte(s1), s2...))
	}
	fmt.Println("append方法:", time.Since(start))
	//Join方法连接
	v := []string{"hello", "world"}
	start = time.Now()
	for i := 0; i < loop; i++ {
		s = strings.Join(v, "")
	}
	fmt.Println("strings.Join方法:", time.Since(start))
	//bytes.writestring方法
	start = time.Now()
	for i := 0; i < loop; i++ {
		var buf bytes.Buffer
		buf.WriteString("hello")
		buf.WriteString("world")
		buf.String()
	}
	fmt.Println("bytes.writestring方法:", time.Since(start))
	//fmt方法连接
	start = time.Now()
	for i := 0; i < loop; i++ {
		s = fmt.Sprintf("%s%s", "hello", "world")
	}
	fmt.Println("fmt方法:", time.Since(start))
	fmt.Println(s)
}
