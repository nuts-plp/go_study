package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct{}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle...")
	fmt.Println("rece from client:msgID:", request.GetMsgID(), ",data:", request.GetData())
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back ping...ping...ping... err:", err)
	}
}

type HelloRouter struct{}

func (p *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle...")
	fmt.Println("rece from client:msgID:", request.GetMsgID(), ",data:", request.GetData())
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Hello...Hello...Hello..."))
	if nil != err {
		fmt.Println("Call back ping...ping...ping... err:", err)
	}
}

func main() {
	//创建服务器对象
	s := znet.NewServer("[Zinx V0.6]")
	//
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Server()
}
