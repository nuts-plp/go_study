package main

import (
	"fmt"
	"net"
)

var myMap = make(map[string]net.Conn)

func send() {
	for {
		var addr, message string
		fmt.Println(&addr)
		fmt.Println(&message)
		n, _ := myMap[addr].Write([]byte(message))
		fmt.Printf("发送消息<%s>成功:%d\n", message, n)
	}
}
func recv(conn net.Conn) {
	for {
		buf := make([]byte, 4096)
		n, _ := conn.Read(buf)
		fmt.Printf("收到来自<%s>的消息:%s\n", conn.RemoteAddr().String(), string(buf[:n]))
	}
	defer conn.Close()
}
func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	for {
		coon, _ := listen.Accept()
		myMap[coon.RemoteAddr().String()] = coon
		fmt.Println(coon.RemoteAddr().String(), "join")
		go recv(coon)
		go send()
	}
}
