package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8848") //监听
	defer listener.Close()                             //延迟关闭监听
	coon, _ := listener.Accept()                       //接收用户

	defer coon.Close() //延迟关闭连接

	buf := make([]byte, 4096) //开辟缓存
	n, _ := coon.Read(buf)    //读取数据
	fmt.Println("收到数据:", string(buf), n)
	coon.Write(buf[:n])
}
