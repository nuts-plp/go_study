package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	v1 "github.com/grpc_gateway/service/v1"
	"google.golang.org/grpc"
)

type Serv struct {
	v1.UnimplementedEchoServiceServer
}

func (s *Serv) Echo(ctx context.Context, message *v1.StrMessage) (*v1.StrMessage, error) {
	log.Printf("recv value:%s\n", message.GetValue())
	return &v1.StrMessage{Value: message.GetValue()}, nil

}

//
//func (s *Serv) EchoBody(ctx context.Context, message *v1.StrMessage) (*v1.StrMessage, error) {
//	return &v1.StrMessage{Value: "I'am " + message.GetValue()}, nil
//
//}
//
//func (s *Serv) EchoFuck(ctx context.Context, message *v1.StrMessage) (*v1.StrMessage, error) {
//	return &v1.StrMessage{Value: "fuck," + message.GetValue()}, nil
//
//}

var grpcServerEnterPoint = flag.String("grpc-server-endpoint", "127.0.0.1:9090", "")
var port = flag.Int("port", 9090, "")

func main() {
	flag.Parse()
	conn, _ := net.Listen("tcp", ":"+strconv.Itoa(*port))
	server := grpc.NewServer()
	v1.RegisterEchoServiceServer(server, &Serv{})
	_ = server.Serve(conn)

}
