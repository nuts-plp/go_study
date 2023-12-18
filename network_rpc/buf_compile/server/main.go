package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/metadata"

	echov1 "github.com/buf_compile/proto/service/nuts/echo/v1"

	"google.golang.org/grpc"
)

var local = flag.String("grpc-server", ":9090", "")

func main() {
	flag.Parse()
	conn, _ := net.Listen("tcp", *local)
	server := grpc.NewServer()
	echov1.RegisterEchoServiceServer(server, &Serv{})
	_ = server.Serve(conn)
}

type Serv struct {
	echov1.UnimplementedEchoServiceServer
}

func (s *Serv) Echo(ctx context.Context, request *echov1.EchoServiceEchoRequest) (*echov1.EchoServiceEchoResponse, error) {
	log.Printf("[server] recv %v\n", request.GetName())
	return &echov1.EchoServiceEchoResponse{Value: "hello! this is 回声桶 ,recv %s\n" + request.GetName()}, nil
}

func (s *Serv) EchoBody(ctx context.Context, req *echov1.EchoServiceEchoBodyRequest) (*echov1.EchoServiceEchoBodyResponse, error) {
	return &echov1.EchoServiceEchoBodyResponse{Value: fmt.Sprintf("[EchoBody] recv id:%v num:%v", req.GetId(), req.GetNum())}, nil
}
func (s *Serv) EchoHeader(ctx context.Context, req *echov1.EchoServiceEchoHeaderRequest) (*echov1.EchoServiceEchoHeaderResponse, error) {
	return &echov1.EchoServiceEchoHeaderResponse{Value: fmt.Sprintf("[EchoHeader] recv id:%v num:%v", req.GetId(), req.GetNum())}, nil

}
func (s *Serv) EchoUpload(stream echov1.EchoService_EchoUploadServer) error {
	//获取传递的文件名称
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return nil
	}
	fileName, ok := md["file_name"]
	if !ok {
		return nil
	}
	file_name := fileName[0]
	log.Printf("[server] recv file name:%v\n", file_name)
	filepath := "file/" + file_name

	//创建存储文件的名称
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("[server] create file failed! file_name:%s\n", file.Name())
		return err
	}
	defer file.Close()
	//接收文件
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&echov1.EchoServiceEchoUploadResponse{Value: fmt.Sprintf("%s was recvived!  ok", file.Name())})
		}
		if err != nil {
			log.Printf("[server] get an unexpected error!  %v\n", err)
			return err
		}
		n, err := file.Write(recv.Content)
		if err != nil {
			log.Printf("[server] content write to file failed! %v\n", err)
			return err
		}
		fmt.Printf("[server] recv %d\n", n)
	}

}
