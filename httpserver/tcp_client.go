package main

import (
	"fmt"
	"log"
	"net"
)

var done chan bool

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	defer conn.Close()

	done = make(chan bool)

	go Hander1(conn)

	<-done
	fmt.Println("客户端程序结束")
}
func Hander1(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		//往服务端写数据
		fmt.Printf("请输入发送的内容：")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容：%s\n", string(buf))
		n, err := conn.Write(buf)
		if err != nil {
			log.Println(err)
			fmt.Println("往服务端写入数据失败")
			return
		}

		if string(buf) == "exit" {
			fmt.Println("客户端关闭退出")
			done <- true
			return
		}

		//从服务端接受
		n1, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("接受数据失败")
			return
		}

		result := buf[:n1]
		fmt.Printf("接收到数据[%d]:%s\n", n, string(result))
	}

}
