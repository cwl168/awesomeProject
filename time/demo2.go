package main

import (
	"fmt"
	"time"
)

//Go中解析RFC3339时间格式
func main() {
	//2006-01-02 15:04:05
	layout := "2006-01-02T15:04:05Z07:00"

	//转化为标准时间格式格式
	toFormatTime := "2018-03-29T09:32:52Z"
	res_time, _ := time.Parse(layout,toFormatTime)
	toFormatTime = res_time.Format("2006-01-02 15:04:05")
	fmt.Println(toFormatTime)


	t, err := time.Parse(layout, "2020-03-29T09:32:52Z")
	fmt.Println(t, err)
}
