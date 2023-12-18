package main

import (
	"context"
	"grpc_stream/proto"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PORT = 9092

func main() {
	//创建链接     不带认证
	conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("dial %d failed! err:%s\n", PORT, err)
		return
	}
	defer conn.Close()
	//创建调用的客户端
	client := proto.NewStreamCClient(conn)
	stream, err := client.ClientStream(context.Background())
	if err != nil {
		log.Printf("streamClient invoke fun to gain an err!!! err:%s\n", err)
		return
	}
	//循环发送消息
	for i := 0; i < 10; i++ {
		//创建发送的消息
		msg := &proto.StreamReq{Request: &proto.StreamMsg{Name: "小潘",
			Value: int32(i)}}
		err := stream.Send(msg)
		if err != nil {
			log.Printf("streamClient send msg failed!  err:%s\n", err)
		}

	}
	recv, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("streamClient client has finished! client gain the closed signal from server failed!  err:%s\n", err)
		return
	}

	log.Printf("streamClient gain the closed msg  name:%s   value:%d\n", recv.GetResponse().Name, recv.GetResponse().Value)
}
