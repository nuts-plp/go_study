package main

import (
	"grpc_stream/proto"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const PORT = 9092

type ServerV struct {
	proto.UnimplementedStreamCServer
}

func (s *ServerV) ClientStream(stream proto.StreamC_ClientStreamServer) error {
	return nil
}

func (s *ServerV) ServerStream(r *proto.StreamReq, stream proto.StreamC_ServerStreamServer) error {
	log.Printf("recv name:%s     value:%d\n", r.GetRequest().Name, r.Request.Value)
	for i := 0; i < 10; i++ {
		respMsg := &proto.StreamResp{Response: &proto.StreamMsg{
			Name:  r.Request.GetName(),
			Value: int32(i) + r.Request.GetValue(),
		}}
		err := stream.Send(respMsg)

		if err != nil {
			log.Printf("")
			return err
		}
	}
	return nil

}
func (s *ServerV) DoubleStream(stream proto.StreamC_DoubleStreamServer) error {
	return nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamCServer(server, &ServerV{})

	conn, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Printf("listen %d failed! err:%s\n", PORT, err)
		return
	}
	defer conn.Close()

	server.Serve(conn)
}
