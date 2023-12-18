package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	servicev1 "github.com/grpc_gateway/proto/service2/nuts/echo/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var grpcEndPoint = flag.String("grpc-server-endpoint", "127.0.0.1:9090", "")

func main() {
	flag.Parse()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := servicev1.RegisterEchoServiceHandlerFromEndpoint(context.Background(), mux, *grpcEndPoint, opts)
	if err != nil {
		log.Printf("register failed!%s\n", err)
		return
	}
	err = http.ListenAndServe(":9091", mux)
	if err != nil {
		log.Printf("listen failed!%s\n", err)
		return
	}

}
