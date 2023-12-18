package main

import (
	"context"
	"fmt"

	echov1 "github.com/buf_test/proto/service/nuts/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dial, _ := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := echov1.NewEchoServiceClient(dial)
	resp, _ := client.Echo(context.Background(), &echov1.Arg{Value: "asdaksda"})
	fmt.Println(resp.GetValue())

}
