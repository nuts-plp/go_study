package main

import (
	"context"
	"log"
	"strconv"

	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/grpc_interceptor/proto"
	"google.golang.org/grpc"
)

const PORT = 9090

func main() {

	conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("[client] connect to %d failed! err:%s\n", err)
		return
	}
	client := pb.NewSearchVClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.Req{Request: "潘丽萍"})
	if err != nil {
		log.Printf("[client] rpc invoke failed! err:%s\n", err)
		return
	}
	log.Printf("[client]  recv value:%s\n", resp.Response)

	stream, err := client.SayHi(context.Background())
	if err != nil {
		log.Printf("[client] stream send failed! err:%s\n", err)
		return
	}
	for i := 0; i < 10; i++ {

		err = stream.Send(&pb.Req{
			Request: "周小林",
		})
		if err != nil {
			log.Printf("[client] recv an err! err:%s\n", err)
			return
		}
	}
	err = stream.CloseSend()
	if err != nil {
		log.Printf("[client] close send failed! err:%s\n", err)
	}

}
