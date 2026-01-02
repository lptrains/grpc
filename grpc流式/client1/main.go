package main

import (
	"client1/proto"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := ":8081"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("%s", fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := proto.NewServiceStreamClient(conn)

	// 正确的流式调用
	stream, err := client.Fun(context.Background(), &proto.Request{Name: "张三"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		response, err := stream.Recv() // 流式对象可调用 Recv()
		if err == io.EOF {
			break // 流结束
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
	}

}
