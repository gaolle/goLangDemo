package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "grpc-go-practice/example/proto"
	"net"
)

const (
	// 服务地址
	Address = "127.0.0.1:50052"
)

// HelloServer is the server API for Hello service.
// type HelloServer interface {
// 	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
// }
// 实现hello.pb.go中的接口
type helloService struct {}

var HelloService = helloService{}

// 实现接口
func (h helloService) SayHello (ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error){
	resp := new(pb.HelloReply)
	resp.Message = "Hello" + in.Name + "."
	return resp, nil
}

func main() {
	// 监听IP
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("listen on " + Address)

	s.Serve(listen)
}