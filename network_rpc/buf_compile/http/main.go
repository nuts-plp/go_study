package main

import (
	"context"
	"flag"
	"fmt"

	echov1 "github.com/buf_compile/proto/service/nuts/echo/v1"

	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var local = flag.String("http-grpc", ":9090", "")

func main() {
	//flag.Parse()
	//mux := runtime.NewServeMux()
	//_ = echo.RegisterEchoServiceHandlerFromEndpoint(context.Background(), mux, *local, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	//_ = http.ListenAndServe(*local, mux)

	flag.Parse()
	mux := runtime.NewServeMux()
	mux.HandlePath("GET", "/ping", handlerPing)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := echov1.RegisterEchoServiceHandlerFromEndpoint(context.Background(), mux, *local, opts)
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

func handlerPing(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	_, _ = fmt.Fprintf(w, "{\"msg\":\"pong\" }")
}

func handlerUpload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "upload failed! err:%d\n", http.StatusBadGateway)
	}
	log.Printf("adasdfsdgasgas\n")
}
