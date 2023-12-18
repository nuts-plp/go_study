package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	echov1 "github.com/buf_test/proto/service/nuts/echo/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var local = flag.String("grpc-server-endpoint", "127.0.0.1:9091", "")

//var grpcEndPoint = flag.String("grpc-server-endpoint", "127.0.0.1:9090", "")

func main() {
	//flag.Parse()
	flag.Parse()
	mux := runtime.NewServeMux()
	fmt.Println(*local)
	//fmt.Println(*grpcEndPoint)
	//mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := echov1.RegisterEchoServiceHandlerFromEndpoint(context.Background(), mux, *local, opts)
	//err := servicev1.RegisterEchoServiceHandlerFromEndpoint(context.Background(), mux, *grpcEndPoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":9091", mux)
	//err = http.ListenAndServe(":9091", mux)
	if err != nil {
		log.Fatal(err)
	}

	//flag.Parse()
	//mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := echov1.RegisterEchoServiceHandlerFromEndpoint(context.Background(), mux, *grpcEndPoint, opts)
	//if err != nil {
	//	log.Printf("register failed!%s\n", err)
	//	return
	//}
	//err = http.ListenAndServe(":9091", mux)
	//if err != nil {
	//	log.Printf("listen failed!%s\n", err)
	//	return
	//}

}
