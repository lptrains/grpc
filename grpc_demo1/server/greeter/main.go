package main

import (
	"context"
	"fmt"
	greeter "greeter/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Hello struct {
	greeter.UnimplementedGreeterServer // 嵌入UnimplementedGreeterServer以实现GreeterServer接口
}

func (this Hello) SayHello(c context.Context, req *greeter.Req) (*greeter.Res, error) {
	fmt.Println(req)
	return &greeter.Res{
		Message: "你好" + req.Name,
	}, nil
}

func main() {

	// 启动服务
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, new(Hello))
	// 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 开始服务
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
