package main

import (
	"context"
	"io"
	"log"
	"net"
	"runtime/debug"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	pb "github.com/grpc_interceptor/proto"
)

const PORT = 9090

type Serv struct {
	pb.UnimplementedSearchVServer
}

func (s *Serv) SayHello(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{Response: "hello!" + req.GetRequest()}, nil

}

func (s *Serv) SayHi(stream pb.SearchV_SayHiServer) error {
	//

	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("[server] client finished its msg")
				break
			} else {
				log.Printf("[server] recv get an err! err:%s\n", err)
				return err
			}
		}
		log.Printf("[server] recv value:%s\n", recv.GetRequest())
	}
	return nil

}

///拦截器
///  go-grpc-middleware 插件实现server使用多个拦截器

// LoggingInterceptor 日志拦截器
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("[server] gRPC Unary method:%s   %v", info.FullMethod, req)
	return handler(ctx, req)

}

// LoggingInterceptorStream  流式
func LoggingInterceptorStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("[server] gRPC stream method:%s", info.FullMethod)
	return handler(srv, stream)

}

// RecoveryInterceptor   recovery拦截器
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			status.Errorf(codes.Internal, "Panic err %s", e)
		}
	}()

	return handler(ctx, req)
}

func main() {
	//tls := GetTLScredentialsByCA()
	opt := []grpc.ServerOption{
		//grpc.Creds(tls),
		grpc_middleware.WithUnaryServerChain(
			LoggingInterceptor,
			RecoveryInterceptor),
		grpc_middleware.WithStreamServerChain(LoggingInterceptorStream),
	}
	//创建服务段
	server := grpc.NewServer(opt...)
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Printf("[server] listen %d failed! err:%s\n", PORT, err)
		return
	}
	//注册
	pb.RegisterSearchVServer(server, &Serv{})
	server.Serve(conn)

}
