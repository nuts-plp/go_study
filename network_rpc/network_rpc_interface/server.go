package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService" //服务名

type HelloServiceInterface = interface { //接口
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error { //操作接口
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
func main() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if nil != err {
		log.Fatal("listen tcp error:", err)
	}
	for {
		conn, err := listener.Accept()
		if nil != err {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
