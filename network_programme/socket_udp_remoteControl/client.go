package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	conn, _ := net.DialUDP("udp", nil, addr)
	var address net.Addr
	go func() {
		for {
			buf := make([]byte, 4096)
			n, addr, _ := conn.ReadFrom(buf)
			address = addr
			fmt.Println("receive<<-", string(buf[:n]), "from", addr.(*net.UDPAddr).String())

		}
	}()
	for {
		var input string
		fmt.Scanln(&input)
		conn.WriteTo([]byte(input), address)
	}

}
