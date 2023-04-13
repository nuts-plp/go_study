package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8887")
	conn, _ := net.ListenUDP("udp", addr) //listenudp 是unconnect连接 即不知道客户端的地址
	defer conn.Close()
	bn := &net.UDPAddr{}
	go func() {
		for {
			buf := make([]byte, 4096)
			n, Addr, _ := conn.ReadFrom(buf)
			//command := exec.Command(string(buf[:n]))
			//output, _ := command.CombinedOutput()
			//fmt.Println(output)
			bn = Addr.(*net.UDPAddr)
			fmt.Println("收到数据<-", string(buf[:n]))
		}
	}()
	for {
		var input string
		fmt.Scanln(&input)
		//n, _ := conn.WriteToUDP([]byte(input), bn)
		n, _ := conn.WriteTo([]byte(input), bn)
		fmt.Println("发送信息->", input, " 成功", "->>", n)

	}

}
