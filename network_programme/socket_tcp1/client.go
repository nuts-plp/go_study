package main

import (
	"fmt"
	"net"
)

func main() {
	coon, _ := net.Dial("tcp", "127.0.0.1:8848")

	defer coon.Close() //关闭

	coon.Write([]byte("hello 潘丽萍！"))

	buf := make([]byte, 4096)
	n, _ := coon.Read(buf)

	fmt.Println("接收数据：", string(buf[:n]), n)

}
