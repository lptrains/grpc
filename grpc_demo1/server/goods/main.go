package main

import (
	"fmt"
	"goods/proto/goods"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Goods struct {
	goods.UnimplementedGoodsServer
}

func (this *Goods) AddGoods(c context.Context, req *goods.AddGoodsReq) (*goods.AddGoodsRes, error) {
	fmt.Println(req)
	return &goods.AddGoodsRes{Message: "数据增加成功", Success: true}, nil
}
func main() {
	//创建服务
	s := grpc.NewServer()
	//注册服务
	goods.RegisterGoodsServer(s, &Goods{})
	//监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听端口失败", err)
	}
	defer listener.Close()
	//启动服务
	s.Serve(listener)
}
