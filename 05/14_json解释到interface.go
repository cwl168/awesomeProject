package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//字符串转[]byte
	b := []byte(`{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666,
    "tttt":4 
}`)

	var t interface{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(t)

	//使用断言判断类型
	m := t.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case bool:
			fmt.Println(k, "is bool", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	//格式化输出
	type Human struct {
		Name string
	}

	var people = Human{Name: "zhangsan"}
	fmt.Println(people)
	fmt.Printf("%v\n", people)
	fmt.Printf("%+v\n", people)
	fmt.Printf("%#v\n", people)
	fmt.Printf("%T\n", people)

	var d rune = 'a'
	var e byte = 255
	var f uint8 = 255
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var data [10]byte
	data[0] = 'A'
	data[1] = 'B'
	var str string = string(data[:])
	fmt.Println(str)

	var str1 string = "helloword"
	var data1 []byte = []byte(str1)
	fmt.Println(data1[0])
}
