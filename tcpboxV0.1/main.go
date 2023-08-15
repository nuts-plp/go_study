package main

import "zinx/tcpnet"

func main() {
	//创建服务器对象
	s := tcpnet.NewServer("[Zinx V0.1]")
	//
	s.Server()
}
