package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:8888")

	defer listen.Close()

	coon, _ := listen.Accept()
	defer coon.Close()

	go func() {

		for {
			buf := make([]byte, 4096)
			n, _ := coon.Read(buf)
			fmt.Println("收到数据:", string(buf[:n]))
		}
	}()

	for {
		var input string
		fmt.Scanln(&input)
		n, _ := coon.Write([]byte(input))
		fmt.Println("发送成功,", n, input)
	}

}
