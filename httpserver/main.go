package main

import (
	"fmt"
	"log"
	"net/http"
)

type user struct {
}

func (u user) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("xxxxx1")
	fmt.Fprint(writer, "<h1>Hello!</h1>")

}

func main() {
	u := user{}
	log.Printf("%v\n", u)
	//ListenAndServe 函数 ,其实是创建了⼀个 http.Server 结构体类型的变量 然后调⽤了这个变量的 ListenAndServe ⽅法
	http.ListenAndServe(":8989", u)

}
