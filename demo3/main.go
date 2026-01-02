package main

import (
	"fmt"
	userservice "iph/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello, world!")
	u := &userservice.User{
		Name:  "张三",
		Email: "123@qq.com",
		Id:    1,
	}
	fmt.Println(u.GetId())
	fmt.Println(u.GetEmail())
	fmt.Println(u.GetName())
	//proto.Maeshal对protobuf数据序列化
	data, _ := proto.Marshal(u)
	fmt.Println(data)

	//反序列化
	u2 := userservice.User{}

	proto.Unmarshal(data, &u2)
	fmt.Printf("%#v\n", u2)
	fmt.Println(u2)
}
