package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc/metadata"

	grpc_tls "github.com/grpc_tls/proto"

	"google.golang.org/grpc"
)

const PORT = 9090

type Serv struct {
	grpc_tls.UnimplementedSearchCServer
}

func (s *Serv) Search(ctx context.Context, req *grpc_tls.Req) (*grpc_tls.Resp, error) {

	return &grpc_tls.Resp{Response: "我喜欢你！" + req.GetRequest()}, nil
}

func TokenAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	mt, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("get metadata failed! \n")
		return nil, nil

	}
	var user string
	var id string
	if i, ok := mt["user"]; ok {
		user = i[0]
	}
	if i, ok := mt["id"]; ok {
		id = i[0]
	}
	if user == "潘丽萍" || id == "123456" {
		log.Printf("pass！\n")
		return handler(ctx, req)
	}
	return nil, errors.New("auth failed！")
}

func main() {
	// 使用拦截器实现token认证
	server := grpc.NewServer(grpc.UnaryInterceptor(TokenAuth))
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Printf("listen to %d failed! %s\n", err)
		return
	}

	grpc_tls.RegisterSearchCServer(server, &Serv{})
	server.Serve(conn)
}
