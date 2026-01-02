package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//建立tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	//简历基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//调用rpc服务方法
	var reply string
	err1 := client.Call("HelloServer.Lph", "hhh", &reply)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println("收到的数据为：", reply)
}
