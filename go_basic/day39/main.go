package main

import (
	"context"
	"fmt"
	servic "go_basic/day39/proto/servi"
	"net"

	"google.golang.org/grpc"
)

type user struct {
	servic.UnimplementedTalkServer
}

func (user *user) SayHello(ctx context.Context, in *servic.Req) (*servic.Resp, error) {
	return &servic.Resp{Response: in.Request + "潘丽萍"}, nil
}
func (u user) SayName(ctx context.Context, in *servic.Req) (*servic.Resp, error) {
	return &servic.Resp{Response: "你好！" + in.Request}, nil
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("监听端口失败！！")
	}
	defer listen.Close()
	server := grpc.NewServer()
	servic.RegisterTalkServer(server, &user{})
	server.Serve(listen)
}
