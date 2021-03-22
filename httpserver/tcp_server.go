package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listen, _ := net.Listen("tcp", ":8080")
	defer listen.Close()
	fmt.Println("等待客户端连接...")
	for {
		con, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go Hander(con)
	}
}
func Hander(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接收数据失败")
			return
		}
		result := buf[:n]
		if string(result) == "exit" {
			fmt.Println("服务端退出连接")
			return
		}
		fmt.Printf("接收到数据:%s\n", string(result))

		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}
