package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	listener, _ := net.ListenTCP("tcp", addr)
	conn, _ := listener.Accept()
	defer conn.Close()
	defer listener.Close()
	go func() {
		for {
			buf := make([]byte, 4096)
			n, _ := conn.Read(buf)
			fmt.Println("receive message<", string(buf[:n]), "> from ", conn.RemoteAddr().String())
		}
	}()
	for {
		var input string
		fmt.Scanln(&input)
		conn.Write([]byte(input))
		fmt.Println("send message<", input, "> to ", conn.RemoteAddr().String())
	}
}
