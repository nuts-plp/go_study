package main

import (
	"net/http"
	"time"

	"github.com/nuts/demo_test/apps/gateway/internal/routers"
)

func main() {
	engine := routers.NewEngine()
	server := http.Server{
		Addr:           ":9090",
		Handler:        engine,
		ReadTimeout:    time.Second * 8,
		WriteTimeout:   time.Second * 8,
		MaxHeaderBytes: 2 << 2,
	}
	_ = server.ListenAndServe()

}
