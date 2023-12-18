package main

import (
	"context"
	"log"
	"strconv"

	grpc_tls "github.com/grpc_tls/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PORT = 9090

type ClientTokenAuth struct {
}

func (receiver ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"user": "潘丽萍",
		"id":   "123456",
	}, nil

}

func (receiver ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}
func main() {
	conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	if err != nil {
		log.Printf("create connection failed! %s", err)
		return
	}
	client := grpc_tls.NewSearchCClient(conn)
	resp, _ := client.Search(context.Background(), &grpc_tls.Req{Request: "潘丽萍"})
	log.Printf("recv value:%s\n", resp.GetResponse())
}
