package main

import (
	"context"
	"fmt"
	service "grpct/proto"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

// Serve 我们自己定义的服务
type Serve struct {
	service.UnimplementedSayHelloServer
}

// token 认证
// SayHi 实现接口，完成我们自己的业务
//func (s *Serve) SayHi(ctx context.Context, request *service.Request) (*service.Response, error) {
//	mt, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		fmt.Printf("get metadata failed！\n")
//		return nil, nil
//	}
//	// 获取元数据token相关数据，和数据库做对比
//	var appid string
//	var appkey string
//	if v, ok := mt["appid"]; ok {
//		appid = v[0]
//	}
//	if v, ok := mt["appkey"]; ok {
//		appkey = v[0]
//	}
//	if appid == "panliping" && appkey == "iloveyou" {
//		return &service.Response{Answer: "HI!" + request.Name}, nil
//	}
//	return nil, errors.New("token data wrong")
//}

func (s *Serve) SayHi(ctx context.Context, request *service.Request) (*service.Response, error) {
	return &service.Response{
		Answer: "hi!" + request.Name,
	}, nil
}
func main() {

	// tls认证
	cert, err := credentials.NewServerTLSFromFile("D:\\JetBrains\\GoLand_workspace\\com.github\\nuts\\go-study\\network_rpc\\grpc_01\\ca.crt", "D:\\JetBrains\\GoLand_workspace\\com.github\\nuts\\go-study\\network_rpc\\grpc_01\\server.keycd ")
	if err != nil {
		fmt.Println("certificate generate failed")
		return
	}
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	// 延迟关闭连接
	defer listen.Close()
	//创建一个grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(cert))
	//在grpc服务中注册我们自己编写的服务
	service.RegisterSayHelloServer(grpcServer, &Serve{})
	//启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("start server failed! err:%s", err)
		return
	}
}
