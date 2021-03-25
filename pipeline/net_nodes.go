package pipeline

import (
	"bufio"
	"net"
)

//开启服务端 ，等待客户端连接
func NetworkSink(addr string, in <-chan int) {
	//监听端口
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer listen.Close()
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		//往这个链接写数据
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSink(writer, in)
	}()
}

//客户端连接
func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		//连接server
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		//从连接读数据
		r := ReaderSource(conn, -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}
