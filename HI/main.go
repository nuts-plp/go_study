package main

import (
	"HI/handler"
	pb "HI/proto/proto"

	"micro.dev/v4/service"
	"micro.dev/v4/service/logger"
)

func main() {
	//Create service
	srv := service.New(
		service.Name("hi"),
	)

	// Register handler
	pb.RegisterHIHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	//ser := grpc.NewService(
	//	micro.Name("hi"),
	//	micro.Version("latest"))
	//ser.Init()
	//pb.RegisterHIHandler(ser.Server())

}
