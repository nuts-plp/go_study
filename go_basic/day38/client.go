package main

import (
	"context"
	"fmt"
	servic "go_basic/day39/proto/servi"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc connect failed")
	}
	defer dial.Close()
	client := servic.NewTalkClient(dial)
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	hello, err := client.SayHello(timeout, &servic.Req{Request: "hello!"})
	if err != nil {
		fmt.Println("rpc failed!")
	}
	fmt.Println(hello.String())
	name, err := client.SayName(timeout, &servic.Req{Request: "潘丽萍"})
	if err != nil {
		fmt.Println("rpc failed")
	}
	fmt.Println(name.String())

}
