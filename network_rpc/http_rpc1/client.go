package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:10001")
	if nil != err {
		panic(err)
	}
	var req float32 //参数
	req = 9
	var resp *float32 //返回值
	client.Call("MathUtils.CalculateCircleArea", req, &resp)
	fmt.Println(*resp)
	var reply *string
	client.Call("MathUtils.Say", "潘丽萍", &reply)
	fmt.Println(*reply)
}
