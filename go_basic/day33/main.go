package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 使用go rpc默认的gob编码格式进行交互
	//{
	//
	//	dial, err := rpc.Dial("tcp", "127.0.0.1:8000")
	//
	//	if err != nil {
	//		fmt.Println("connect to 127.0.0.1:8000 failed!")
	//		return
	//	}
	//	var str string
	//	start := time.Now().Unix()
	//	err = dial.Call("stu.SayHi", "潘丽萍", &str)
	//	if err != nil {
	//		fmt.Println("rpc failed!")
	//		return
	//	}
	//	end := time.Now().Unix()
	//	fmt.Println(end - start)
	//	fmt.Println(str)
	//}

	//使用jsonrpc编码格式进行交互
	{
		dial, err := net.Dial("tcp", "47.92.232.226:8000")
		if err != nil {
			fmt.Println("connect failed!")
			return
		}
		// 使用json进行编码解码
		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(dial))
		var str string
		err = client.Call("stu.SayHi", "潘丽萍", &str)
		if err != nil {
			fmt.Println("RPC failed!")
			return
		}
		fmt.Println(str)
	}
	//dial.Go("stu.SayHi", "周小林", &str, done)
}
