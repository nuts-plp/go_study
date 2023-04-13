package main

import (
	"fmt"
	"zinx/tcpiface"
	"zinx/tcpnet"
)

type PingRouter struct{}

func (p *PingRouter) PreHandle(request tcpiface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back before ping err:", err)
	}
}
func (p *PingRouter) Handle(request tcpiface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back ping...ping...ping... err:", err)
	}
}
func (p *PingRouter) PostHandle(request tcpiface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if nil != err {
		fmt.Println("Call back after ping err:", err)
	}
}

func main() {
	//创建服务器对象
	s := tcpnet.NewServer("[Zinx V0.2]")
	//
	s.AddRouter(&PingRouter{})
	s.Server()
}
