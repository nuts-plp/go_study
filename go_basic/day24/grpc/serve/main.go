package main

import (
	"net"
	"net/http"
	"net/rpc"
)

const (
	address = "127.0.0.1:8090"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
type Po struct{}

func (p *Po) SayHello(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (p *Po) SayHi(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
func main() {
	rpc.Register(new(Po))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	http.Serve(listen, nil)
}
