package main

import (
	"fmt"
	"net"
)

func main() {
	dial, _ := net.Dial("tcp", "127.0.0.1:8888")
	defer dial.Close()

	go func() {
		for {
			buf := make([]byte, 4096)
			n, _ := dial.Read(buf)
			fmt.Println("成功接收:", string(buf[:n]))
		}
	}()

	for {
		var input string
		fmt.Scanln(&input)
		n, _ := dial.Write([]byte(input))
		fmt.Println("发送成功:", n, input)
	}
}
