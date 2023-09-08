package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Stu struct {
}

func (s Stu) SayHi(para string, para2 *string) error {
	*para2 = "hello!" + para
	return nil

}
func main() {
	err := rpc.RegisterName("stu", new(Stu))
	if err != nil {
		fmt.Println("unknown reason!")
		return
	}
	listen, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("net listening failed!")
		return
	}
	defer listen.Close()
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("create connection failed!")
			continue
		}
		// go的rpc默认使用gob编码格式  此编码格式是go独有的
		//go rpc.ServeConn(accept)

		// 使用jsonrpc格式编码代替gob编码格式      jsonrpc编码格式不支持http协议
		go rpc.ServeCodec(jsonrpc.NewServerCodec(accept))
	}
}
