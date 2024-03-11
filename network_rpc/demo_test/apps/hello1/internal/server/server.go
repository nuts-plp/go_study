package server

import (
	"context"
	"flag"
	"fmt"
	"net"

	"go.etcd.io/etcd/client/v3/naming/endpoints"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	hello_servicev1 "github.com/nuts/demo_test/idl/service/nuts/hello/v1"
)

var (
	localFlag = flag.String("hello1", ":9091", "hello1 service endpoint")
	etcdFlag  = flag.String("etcd", ":2379", "etcd server endpoint")
)

type Serv struct {
	*hello_servicev1.UnimplementedHelloServiceServer
}

func (s *Serv) Echo(ctx context.Context, request *hello_servicev1.EchoRequest) (response *hello_servicev1.EchoResponse, err error) {
	fmt.Println(request.GetValue() + ",潘丽萍")
	return &hello_servicev1.EchoResponse{Value: request.Value + ",潘丽萍"}, nil
}

func (s *Serv) Run() {
	flag.Parse()
	listen, err := net.Listen("tcp", *localFlag)
	if err != nil {
		fmt.Printf("[server1] listen %s failed! %v\n", *localFlag, err)
		return
	}
	// 创建一个grpc服务
	server := grpc.NewServer()
	//将grpc服务与此服务实例绑定
	hello_servicev1.RegisterHelloServiceServer(server, &Serv{})
	// 创建一个etcd客户端
	client, err := clientv3.NewFromURL(*etcdFlag)
	if err != nil {
		fmt.Printf("[server1] create etcd client failed! %s\n", err)
	}
	//通过etcd客户端实现一个管理，并将此服务注册到etcd
	manager, err := endpoints.NewManager(client, "nuts/hello/v1")
	if err != nil {
		fmt.Printf("[server1] create manager failed! %v\n", err)
		return
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	lease, err := client.Grant(ctx, 2004)
	if err != nil {
		fmt.Printf("[server1] create lease failed! %v\n", err)
		return
	}
	//将此服务注册到注册中心
	err = manager.AddEndpoint(ctx, "nuts/hello/v1/hello1", endpoints.Endpoint{Addr: *localFlag}, clientv3.WithLease(lease.ID))
	if err != nil {
		fmt.Printf("[server1] service register failed! %v\n", err)
		return
	}
	// 监听服务
	server.Serve(listen)
}
func NewServer() *Serv {
	return &Serv{}
}
