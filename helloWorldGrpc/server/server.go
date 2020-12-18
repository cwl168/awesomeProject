package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	pb "learnDemo/helloWorldGrpc/proto"
	"log"
	"net"
)

// 定义server服务
type server struct {
}

// 实现server服务接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//利用函数 FromIncomingContext从context中获取metadata:
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	//time.Sleep(2*time.Second)
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "i am key" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	log.Printf("Received: %v.\nToken info: appid=%s,appkey=%s", in.GetName(), appid, appkey)

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

	reflection.Register(s)
	// 注册server服务
	pb.RegisterGreeterServer(s, &server{})

	log.Println("Listen on 127.0.0.1:8888...")
	s.Serve(listen)
}

// 查询服务列表 grpcurl -plaintext  localhost:8888 list
// 查询服务提供的方法 grpcurl -plaintext  localhost:8888 list helloworld.Greeter
// 查看更详细的描述  grpcurl -plaintext  localhost:8888 describe helloworld.Greeter
// 调用服务的方法  grpcurl -plaintext  localhost:8888 helloworld.Greeter.SayHello
