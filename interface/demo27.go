package main

import "fmt"

//1.定义一个结构体，实现error接口
type RectError struct {
	//字段
	msg string  //错误的文字描述
	wid float64 //
	len float64 //
}

func (r *RectError) Error() string {
	return fmt.Sprintf("宽度：%.2f,高度：%.2f,错误提示：%s", r.wid, r.len, r.msg)
}

func rectArea(width, length float64) (float64, error) { // -4, 8
	// 1.拼接错误的信息
	errMsg := "" // string
	if width <= 0 {
		errMsg = "宽度不能小于等于零"
	}
	if length <= 0 {
		if errMsg == "" {
			errMsg = "长度不能小于等于零"
		} else {
			errMsg += ",长度也不能小于等于零"
		}
	}
	//2. 判断错误的信息：""
	if errMsg != "" {
		// 返回结果
		return 0, &RectError{errMsg, width, length}
	}
	//3,如果没有错误，进行计算
	area := width * length
	return area, nil

}
func main() {
	/*
	   自定义错误的类型：
	      error-->interface{Error() string}

	      只要实现了该接口，就可以作为错误的一种类型
	*/
	area, err := rectArea(-4, -8)
	if err != nil { //有错误，打印错误信息
		fmt.Println(err)
	} else { //没有错误，打印面积
		fmt.Println(area)
	}

}
