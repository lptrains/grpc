package main

import (
	greeter "Lph/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到服务器
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	client := greeter.NewGreeterClient(conn)
	// 调用服务
	res, err := client.SayHello(context.Background(), &greeter.Req{
		Name: "张三",
	}) //context.Background()用于main函数初始化和测试
	if err != nil {
		fmt.Println("调用失败", err)
		return
	}
	fmt.Println("调用成功", res.Message)

}
