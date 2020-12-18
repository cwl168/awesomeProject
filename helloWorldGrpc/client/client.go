package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "learnDemo/helloWorldGrpc/proto"
	"log"
	"time"
)

const (
	timestampFormat = time.StampNano
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

//是否需要基于TLS认证进行安全传输
func (c customCredential) RequireTransportSecurity() bool {
	return false
}
func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	// 连接grpc服务
	conn, err := grpc.Dial(":8888", opts...)
	if err != nil {
		log.Fatal(err)
	}
	// 很关键
	defer conn.Close()

	// 初始化客户端
	c := pb.NewGreeterClient(conn)
	// 新建一个有 metadata 的 context  发送metadata
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	//1秒超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	newCtx := metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	// 发起请求
	response, err := c.SayHello(newCtx, &pb.HelloRequest{Name: "BaoBao"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Message)
}
