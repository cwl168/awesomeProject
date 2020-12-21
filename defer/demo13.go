package main

import (
	"fmt"
	"log"
	"net/http"
)

//过早关闭http响应体
func main() {
	resp, err := http.Get("baidu.com")
	//1
	//defer resp.Body.Close()
	if err != nil {
		log.Fatalf("http get err is : %s \n", err)
	}
	//2
	defer resp.Body.Close()
	fmt.Println(resp)
}
