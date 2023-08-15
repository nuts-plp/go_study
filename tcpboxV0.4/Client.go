package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if nil != err {
		fmt.Printf("dail err:", err)
		return
	}
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		//var in string
		//n, err := fmt.Scan(&in)
		//if nil != err {
		//	fmt.Println("Scan message err:", err)
		//	continue
		//}

		n, err := conn.Write([]byte("你好！"))
		if nil != err {
			fmt.Println("write err:", err)
			return
		}
		fmt.Println("Server sent ", "你好啊！")
		n, err = conn.Read(buf[:n])
		if nil != err {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("Server read ", string(buf))
		time.Sleep(5 * time.Second)
	}

}
