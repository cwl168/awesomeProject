package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type M struct {
	Data map[string]interface{} `json:"data"`
}

func main() {

	g := new(M)
	g.Data = make(map[string]interface{})
	g.Data["key"] = "123"
	s, _ := json.Marshal(g)
	fmt.Println(string(s))

	address := "www.baidu.com:8080"
	//TCPAddr 是一个包含 IP和port的 struct
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tcpAddr:%+v\n", tcpAddr)

}
