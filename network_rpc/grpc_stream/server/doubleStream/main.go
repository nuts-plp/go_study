package main

import (
	"grpc_stream/proto"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const PORT = 9092

type ServeC struct {
	proto.UnimplementedStreamCServer
}

func (receiver *ServeC) ClientStream(stream proto.StreamC_ClientStreamServer) error {
	return nil

}
func (receiver *ServeC) ServerStream(req *proto.StreamReq, stream proto.StreamC_ServerStreamServer) error {
	return nil

}

func (receiver *ServeC) DoubleStream(stream proto.StreamC_DoubleStreamServer) error {
	for i := 0; i < 100; i++ {
		err := stream.Send(&proto.StreamResp{
			Response: &proto.StreamMsg{
				Name:  "潘丽萍",
				Value: int32(i),
			},
		})
		if err != nil {
			log.Printf("streamDouble get an err when sending! err:%s\n", err)
			break
		}
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Printf("streamDouble client recv finished")
			break
		}
		if err != nil {
			log.Printf("streamDoouble [client] gvet an err !   err:%s\n", err)
			break
		}
		log.Printf("streamDouble  client recv name:%s   value:%d\n", recv.Request.Name, recv.Request.Value)
	}
	return nil

}

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamCServer(server, &ServeC{})

	conn, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Printf("'listen %d failed! err:%s\n", PORT, err)
		return
	}
	defer conn.Close()
	server.Serve(conn)
}
