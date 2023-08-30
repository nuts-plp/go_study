package main

import (
	"context"

	"github.com/micro/go-micro"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Hi struct{}

func (receiver Hi) GetHi(ctx context.Context, req *Request, res *Response) error {
	res.Message = "hi!" + req.Name
	return nil
}

func main() {

	ser := micro.NewService(
		micro.Name("HI"),
		micro.Version("v1.0"),
		micro.Address(":8090"),
	)
	ser.Init()
	micro.RegisterHandler(ser.Server(), new(Hi))
	ser.Run()
	/////
}
