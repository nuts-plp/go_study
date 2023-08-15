package main

import (
	"fmt"
	"net"
)

func process(coon net.Conn) {
	defer coon.Close()
	buf := make([]byte, 1024)
	for {
		n, err := coon.Read([]byte(buf)) //读取内容到buf中
		if err != nil {
			fmt.Printf("%v 远程客户端断开,err:%v", coon.RemoteAddr().String(), err)
			return
		}
		fmt.Print(string(buf[:n])) //打印信息到终端
	}

}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败！")
		return
	}

	defer listen.Close()
	for {
		fmt.Println("等待连接。。。。。")
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("获取连接失败！")
			return
		}
		fmt.Printf("accept() success con:%v ip:%v", coon, coon.RemoteAddr().String())
		fmt.Println("连接成功！")
		go process(coon)
	}
}
