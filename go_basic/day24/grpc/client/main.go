package main

import (
	"fmt"
	"net/rpc"
)

const (
	address = "127.0.0.1:8090"
)

type Args struct {
	A, B int
}
type reply int

func main() {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		return
	}
	args := Args{8, 6}
	var h reply
	err = client.Call("Po.SayHello", args, &h)
	if err != nil {
		return
	}
	fmt.Println(h, "|||||")
	call := client.Go("Po.SayHi", args, &h, nil)
	j := <-call.Done
	fmt.Println(h, "////////", j)
}
