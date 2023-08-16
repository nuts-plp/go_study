package main

import (
	"fmt"
	"net"
)

func main() {
	dial, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	udp, _ := net.DialUDP("udp", nil, dial)
	fmt.Println(udp, dial)
	write, err := udp.WriteTo([]byte("afasdfasdf"), dial)
	fmt.Println(write, err)
}
