package main

import (
	"fmt"
	"net"
	"os/exec"
)

var myMap = make(map[string]net.Conn)

func execFunc(coon net.Conn) {
	for {
		buf := make([]byte, 4096)
		n, _ := coon.Read(buf)
		command := exec.Command(string(buf[:n]))
		output, _ := command.CombinedOutput()
		n, _ = coon.Write(output)
		fmt.Println("指令：", string(buf[:]), "执行成功：", n)

	}
	defer coon.Close()
}
func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:8888")
	defer listen.Close()
	for {
		coon, _ := listen.Accept()
		fmt.Println(coon.RemoteAddr().String(), "join")
		go execFunc(coon)
	}
}
