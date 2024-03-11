package main

import "github.com/nuts/demo_test/apps/hello2/internal/server"

func main() {
	serv := server.NewServer()
	serv.Run()
}
