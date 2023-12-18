package main

import (
	"context"
	"fmt"
	service "grpct/proto"
	"time"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

type ClientTokenAuth struct{}

func (receiver ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "panliping",
		"appkey": "iloveyou",
	}, nil
}

func (receiver ClientTokenAuth) RequireTransportSecurity() bool {
	return false

}

func main() {
	cert, _ := credentials.NewClientTLSFromFile("D:\\JetBrains\\GoLand_workspace\\com.github\\nuts\\go-study\\network_rpc\\grpc_01\\ca.crt", "*.nuts.com")
	//创建链接   禁用安全传输，没有加密和验证

	//使用token认证
	//conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()),
	//	grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(cert))
	if err != nil {
		fmt.Printf("conn create failed! err:%s\n", err)
		return
	}
	defer conn.Close()
	// 创建一个客户端
	start := time.Now()
	client := service.NewSayHelloClient(conn)
	//调用方法，获取响应
	resp, err := client.SayHi(context.Background(), &service.Request{Name: "潘丽萍"})
	if err != nil {
		fmt.Printf("resp failed! err:%s\n", err)
		return
	}
	sub := time.Now().Sub(start)
	//打印响应
	fmt.Println(resp.GetAnswer())
	fmt.Println(sub)
}
