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

type ServerV struct {
	proto.UnimplementedStreamCServer
}

func (s *ServerV) ClientStream(stream proto.StreamC_ClientStreamServer) error {
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Printf("server recv signal to close!\n")
			return stream.SendAndClose(&proto.StreamResp{Response: &proto.StreamMsg{
				Name:  "streamClient has received the signal of close,closing!!!!",
				Value: int32(0)}})
		}
		if err != nil {
			log.Printf("streamClient server recv failed! err:%s\n", err)
			return err
		}
		log.Printf("streamClient recv name:%s   value:%d   \n", recv.GetRequest().Name, recv.GetRequest().Value)

	}

}

func (s *ServerV) ServerStream(req *proto.StreamReq, stream proto.StreamC_ServerStreamServer) error {
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
		log.Printf("listen %d faild! err:%s\n", PORT, err)
		return
	}
	defer conn.Close()
	err = server.Serve(conn)
	if err != nil {
		panic("server serve failed!！")
	}
}
