package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:8888")

	defer listen.Close()
	accept, _ := listen.Accept()

	defer accept.Close()
	for {
		buf := make([]byte, 4096)
		n, _ := accept.Read(buf)
		command := exec.Command(string(buf[:n]))
		output, _ := command.CombinedOutput()

		write, _ := accept.Write(output)
		fmt.Printf("收到命令:%s-%d\n", command, write)

	}

}
