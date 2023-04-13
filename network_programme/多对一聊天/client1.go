package main

import (
	"fmt"
	"net"
)

func main() {
	coon, _ := net.Dial("tcp", "127.0.0.1:8889")
	defer coon.Close()
	go func() {
		for {
			buf := make([]byte, 4096)
			n, _ := coon.Read(buf)
			fmt.Println("收到消息:", string(buf[:n]), n)
		}
	}()
	for {
		var input string
		fmt.Scanln(&input)
		n, _ := coon.Write([]byte(input))
		fmt.Println("发送成功！", n)
	}

}
