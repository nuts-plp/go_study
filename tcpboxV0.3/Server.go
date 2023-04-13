package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct{}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back before ping err:", err)
	}
}
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back ping...ping...ping... err:", err)
	}
}
func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back after ping err:", err)
	}
}

func main() {
	//创建服务器对象
	s := znet.NewServer("[Zinx V0.3]")
	//
	s.AddRouter(&PingRouter{})
	s.Server()
}
