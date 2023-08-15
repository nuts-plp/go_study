package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	coon, _ := net.Dial("tcp", "127.0.0.1:8888")
	reader := bufio.NewReader(os.Stdin)
	var n int
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端是io.EOF错误")
			}
			return
		}
		stdinString := strings.Trim(readString, " \r\n")
		if stdinString == "quit" || stdinString == "exit" {
			fmt.Println("程序退出")
			return
		}
		n, err = coon.Write([]byte(readString))
		if err != nil {
			fmt.Println("写入错误")
			return
		}
		fmt.Printf("发送了%d字节 \n", n)

	}

}
