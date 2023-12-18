package main

import (
	"context"
	"log"
	"strconv"

	grpc_tls "github.com/grpc_tls/proto"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

const PORT = 9091

func main() {
	//客户端生成证书
	cert, err := credentials.NewClientTLSFromFile("./../../conf/tls/server.pem", "*.bytefree.com")
	if err != nil {
		log.Printf("[client] generate cert failed! err:%s\n", err)
		return
	}
	//创建连接
	conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(cert))
	if err != nil {
		log.Printf("[client] connect to %d failed! err:%s\n", PORT, err)
		return
	}
	defer conn.Close()
	client := grpc_tls.NewSearchCClient(conn)
	resp, err := client.Search(context.Background(), &grpc_tls.Req{Request: "潘丽萍"})
	if err != nil {
		log.Printf("[client] rpc invoke failed! err:%s\n", err)
		return
	}
	log.Printf("[client] resp  value:%s\n", resp.GetResponse())

}
