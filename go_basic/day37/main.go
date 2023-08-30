package main

import (
	"fmt"
	ppc "go_basic/day37/proto/service"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := &ppc.Reqc{
		Name: "dasda",
	}
	marshal, _ := proto.Marshal(req)
	fmt.Println(marshal)
	vv := &ppc.Reqc{}

	_ = proto.Unmarshal(marshal, vv)
	fmt.Println(vv)
	fmt.Println(vv.String())

}
