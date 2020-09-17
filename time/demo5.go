package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now() //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	fmt.Println(t.Format("20060102150405"))

	//当前时间戳
	t1 := time.Now().Unix() //1564552562
	fmt.Println(t1)
	//时间戳转化为具体时间
	fmt.Println(time.Unix(t1, 0).String())

	//基本格式化的时间表示
	fmt.Println(time.Now().String()) //2019-07-31 13:56:35.7766729 +0800 CST m=+0.005042501

	fmt.Println(time.Now().Format("2006-01-02"))          //2019-07-31
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2019-07-31 13:57:52

	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local))
	fmt.Println(time.Now().Format(time.RFC3339))

	the_time, _ := time.Parse("2006-01-02 15:04:05", "2020-09-01 09:04:41")
	fmt.Println(int(time.Since(the_time).Hours() / 24))
}
