package logic

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	msg2 "github.com/nuts/demo_test/pkg/msg"

	hello_servicev1 "github.com/nuts/demo_test/idl/service/nuts/hello/v1"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
)

var (
	etcdFlag = flag.String("etcd", ":2379", "etcd endpoints")
)

func Login(ctx *gin.Context) {

}
func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func Hello(ctx *gin.Context) {
	//解析ctx中的参数
	msg := new(msg2.Req)
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		fmt.Printf("parse para failed! %v\n", err)
		return
	}
	flag.Parse()
	url, err := clientv3.NewFromURL(*etcdFlag)
	if err != nil {
		fmt.Printf("create etcd client failed! %v\n", err)
		return
	}
	builder, err := resolver.NewBuilder(url)
	if err != nil {
		fmt.Printf("create builder failed! %v\n", err)
		return
	}
	//创建一个grpc客户端
	dial, err := grpc.Dial("etcd:///nuts/hello/v1",
		grpc.WithResolvers(builder),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("connect to etcd failed! %v\n", err)
		return
	}
	//
	client := hello_servicev1.NewHelloServiceClient(dial)
	echo, err := client.Echo(context.TODO(), &hello_servicev1.EchoRequest{
		Value: msg.Value,
	})
	if err != nil {
		fmt.Printf("rpc failed! %v\n", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"resp": echo.GetValue()})
}
