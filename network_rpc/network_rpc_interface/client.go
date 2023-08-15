package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//const HelloServiceName = "path/to/pkg.HelloService" //服务名
//
//type HelloServiceInterface = interface { //接口
//	Hello(request string, reply *string) error
//}
type HelloServiceClient struct {
	*rpc.Client
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

func DailHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if nil != err {
		return nil, err
	}
	return &HelloServiceClient{c}, err

}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func main() {
	client, err := DailHelloService("tcp", ":1234")
	if nil != err {
		log.Fatal("dailing error:", err)
	}
	var reply string
	err = client.Hello("golang", &reply)
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(reply)

}
