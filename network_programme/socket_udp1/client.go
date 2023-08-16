package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8887")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	go func() {
		for {
			buf := make([]byte, 4096)
			n, _ := conn.Read(buf)

			fmt.Println("收到来自<", conn.RemoteAddr().String(), ">的命令:", string(buf[:n]))
		}
	}()
	for {
		var input string
		fmt.Scanln(&input)
		conn.Write([]byte(input))
		fmt.Println("发送数据:", input)

	}
}
