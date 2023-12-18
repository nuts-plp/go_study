package main

import (
	"context"
	"grpc_stream/proto"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PORT = 9092

func main() {
	conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("dial %d failed! err:%s\n", PORT, err)
		return
	}
	defer conn.Close()
	client := proto.NewStreamCClient(conn)
	stream, err := client.DoubleStream(context.Background())
	if err != nil {
		log.Printf("streamDouble  client creates stream failed! err:%s\n", err)
		return
	}
	for i := 0; i < 100; i++ {
		err := stream.Send(&proto.StreamReq{
			Request: &proto.StreamMsg{Name: "周小林",
				Value: int32(i)},
		})
		if err != nil {
			log.Printf("")
		}
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Printf("streamDouble recv finished")
			break
		}
		if err != nil {
			log.Printf("streamDouble recv an err ! err:%s\n", err)
			return
		}
		log.Printf("streamDouble recv name:%s   value:%d \n", recv.Response.Name, recv.Response.Value)
	}
	//客户端发起关闭
	_ = stream.CloseSend()
}
