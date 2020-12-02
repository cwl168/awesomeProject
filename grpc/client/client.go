package main

import (
	"context"
	"google.golang.org/grpc"
	"learnDemo/gRPC/proto"
	"log"
)

func main() {
	// 连接grpc服务
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// 很关键
	defer conn.Close()

	// 初始化客户端
	c := hellworld.NewLoveClient(conn)

	// 发起请求
	response, err := c.SayHello(context.Background(), &hellworld.HelloRequest{Name: "BaoBao"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Result)
}
