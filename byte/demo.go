package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type s1 struct {
	a int
	s []string
}
type s2 struct {
	a int
}
type error1 interface {
	Error() string
}

func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10) //以10进制方式追加
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str)) //4567false"abcdefg"'单'

	var reader io.Reader

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
	}

	reader = tty
	fmt.Println(reader)

	//var x = []int{1, 2, 3}
	//fmt.Println(x == x)

	//var x interface{} = []int{1, 2, 3}
	//fmt.Println(x == x)

	//var w io.Writer
	//fmt.Printf("w==nil:%t\n", w == nil)
	//w = os.Stdout
	//fmt.Printf("w==nil:%t\n", w == nil)
	//w = new(bufio.Writer)
	//fmt.Printf("w==nil:%t\n", w == nil)
	//w = nil
	//fmt.Printf("w==nil:%t\n", w == nil)

	//动态类型与静态类型
	var i interface{}

	i = 18
	i = "Go编程时光"

	fmt.Println(i)

	fmt.Println(reflect.DeepEqual(s1{1, []string{"aa", "bb"}}, s1{1, []string{"aa", "bb"}}))
	var ch1, ch2, ch3 byte
	ch1 = 'a'  //字符赋值
	ch2 = 65   //字符的ascii码赋值
	ch3 = '\n' //转义字符
	fmt.Printf("ch1 = %c, ch2 = %c, %c", ch1, ch2, ch3)

	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	fmt.Println(a, b, c)

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a1 := array[:3]
	fmt.Println(a1)

}
