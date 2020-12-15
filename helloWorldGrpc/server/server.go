package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learnDemo/helloWorldGrpc/proto"
	"log"
	"net"
)

// 定义server服务
type server struct {
}

// 实现server服务接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 监听8888端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	// 实例化grpc server
	s := grpc.NewServer()

	// 注册server服务
	pb.RegisterGreeterServer(s, &server{})

	log.Println("Listen on 127.0.0.1:8888...")
	s.Serve(listen)
}
