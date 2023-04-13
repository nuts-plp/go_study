package main

import (
	"fmt"
	"net"
)

func main() {
	coon, _ := net.Dial("tcp", "127.0.0.1:8888")
	defer coon.Close()

	for {
		var input string
		fmt.Scanln(&input)
		n, _ := coon.Write([]byte(input))
		fmt.Printf("执行指令:%s-%d\n", input, n)
		buf := make([]byte, 4096)
		n, _ = coon.Read(buf)
		fmt.Printf("指令结果:%d\n %s", n, string(buf[:n]))
	}
}
