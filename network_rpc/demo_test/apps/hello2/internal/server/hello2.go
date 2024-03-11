package server

import (
	"context"
	"flag"
	"fmt"
	"net"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"

	"google.golang.org/grpc"

	hello_servicev1 "github.com/nuts/demo_test/idl/service/nuts/hello/v1"
)

var (
	localFlog = flag.String("[server2]", ":9092", "local service2")
	etcdFlog  = flag.String("[serve2]", ":2379", "etcd server endpoints")
)

type Serv struct {
	*hello_servicev1.UnimplementedHelloServiceServer
}

func NewServer() *Serv {
	return &Serv{}
}

func (s *Serv) Echo(ctx context.Context, request *hello_servicev1.EchoRequest) (*hello_servicev1.EchoResponse, error) {
	fmt.Printf("%v,周小林\n", request.GetValue())
	return &hello_servicev1.EchoResponse{Value: request.GetValue() + ",周小林"}, nil

}
func (s Serv) Run() {
	flag.Parse()
	conn, err := net.Listen("tcp", *localFlog)
	if err != nil {
		fmt.Printf("[server2] listen %s failed! %v\n", *localFlog, err)
		return
	}
	//创建grpc服务并绑定服务
	server := grpc.NewServer()
	hello_servicev1.RegisterHelloServiceServer(server, &Serv{})
	//创建一个etcd客户端
	client, err := clientv3.NewFromURL(*etcdFlog)
	if err != nil {
		fmt.Printf("[server2] create etcd client failed! %v\n", err)
		return
	}
	manager, err := endpoints.NewManager(client, "nuts/hello/v1")
	if err != nil {
		fmt.Printf("[server2] create manager failed! %v\n", err)
		return
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	lease, err := client.Grant(ctx, 2000)
	if err != nil {
		fmt.Printf("[server2] create lease failed! %v\n", err)
		return
	}
	err = manager.AddEndpoint(ctx, "nuts/hello/v1/hello2", endpoints.Endpoint{Addr: *localFlog}, clientv3.WithLease(lease.ID))
	if err != nil {
		fmt.Printf("[server2] service register failed! %v\n", err)
		return
	}
	server.Serve(conn)
}
