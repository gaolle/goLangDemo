package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "grpc-go-practice/example/proto"
)

const (
	// 服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	connect, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln("connect:", err)
	}
	// 延迟关闭
	defer connect.Close()
	// 初始化客户端
	c := pb.NewHelloClient(connect)
	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println("reply", r.Message)
}