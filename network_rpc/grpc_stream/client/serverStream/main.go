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
	stream, err := client.ServerStream(context.Background(), &proto.StreamReq{Request: &proto.StreamMsg{
		Name: "潘丽萍",
	}})
	if err != nil {
		log.Printf("")
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("recv failed ! err:%s\n")
			return
		}
		log.Printf("recv name:%s   value:%d    \n", msg.GetResponse().Name, msg.GetResponse().Value)
		log.Printf("recv name:%s   value:%d    \n", msg.Response.Name, msg.Response.Value)

	}
}
