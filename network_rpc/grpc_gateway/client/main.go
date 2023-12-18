package main

import (
	"context"
	"fmt"
	"strconv"

	v1 "github.com/grpc_gateway/service/v1"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

const PORT = 9090

func main() {
	conn, _ := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := v1.NewEchoServiceClient(conn)
	resp, _ := client.Echo(context.Background(), &v1.StrMessage{Value: "潘丽萍"})
	fmt.Println("recv  ", resp.GetValue())
}
