package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc/credentials"

	grpc_tls "github.com/grpc_tls/proto"
	"google.golang.org/grpc"
)

const PORT = 9091

type Serv struct {
	grpc_tls.UnimplementedSearchCServer
}

func (s *Serv) Search(ctx context.Context, req *grpc_tls.Req) (*grpc_tls.Resp, error) {
	return &grpc_tls.Resp{Response: "hello!" + req.Request + ",没想到吧，不是search而是hello！"}, nil
}
func main() {
	// 加载证书文件，生成tls凭证
	cert, err := credentials.NewServerTLSFromFile("./../../conf/tls/server.pem", "./../../conf/tls/server.key")
	if err != nil {
		log.Printf("[server] generate cert faield! err:%s\n", err)
		return
	}
	//将证书添加到服务端中
	server := grpc.NewServer(grpc.Creds(cert))
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Printf("[server] listen %d failed! err:%s\n", err)
		return
	}
	defer conn.Close()
	grpc_tls.RegisterSearchCServer(server, &Serv{})

	server.Serve(conn)

}
