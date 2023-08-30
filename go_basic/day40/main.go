package main

import (
	"net"
	"net/http"
)

func main() {
	listen, _ := net.Listen("tcp", ":8090")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello!,潘丽萍"))
	})
	_ = http.Serve(listen, nil)
}
