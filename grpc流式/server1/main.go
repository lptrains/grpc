package main

import (
	"fmt"
	"log"
	"net"
	"server1/proto"

	"google.golang.org/grpc"
)

type ServiceStream struct {
	proto.UnimplementedServiceStreamServer //满足接口规范
}

func (ServiceStream) Fun(Request *proto.Request, stream proto.ServiceStream_FunServer) error {
	fmt.Println(Request)
	for i := 0; i < 10; i++ {
		stream.Send(&proto.Response{
			Text: fmt.Sprintf("第%d轮数据", i),
		})

	}
	return nil
}
func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterServiceStreamServer(server, &ServiceStream{})

	server.Serve(listen)

}
