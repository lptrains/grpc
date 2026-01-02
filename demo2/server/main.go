package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义类对象
type World struct{}

// 绑定类方法
func (this *World) Lph(req string, res *string) error {
	*res = req + "哦原来你也弹钢琴"
	return nil
}

func main() {
	//注册rpc服务 维护一个hash表，key值是服务名称，value值是服务的地址
	rpc.RegisterName("HelloServer", new(World))
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	//设置监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//给当前连接提供针对json格式的rpc服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
