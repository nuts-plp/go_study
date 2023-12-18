package main

import (
	"context"
	"flag"
	"fmt"

	echov1 "github.com/buf_compile/proto/service/nuts/Echo/v1"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

var local = flag.String("client-grpc", ":9090", "")

func main() {
	conn, _ := grpc.Dial(*local, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := echov1.NewEchoServiceClient(conn)
	echo, _ := client.Echo(context.Background(), &echov1.EchoServiceEchoRequest{Name: "潘丽萍"})
	fmt.Println(echo.GetValue())

}
