package main

import (
	"bufio"
	"fmt"
	"net"
	"io"
)
/*
	1、监听端口
	2、建立连接
	3、接收信息
	4、处理信息
	5、关闭连接	
*/
//ReceiveMessage 每接收信息开启一个goroutine处理


func main(){
	//监听端口
	listen,err :=net.Listen("tcp","127.0.0.1:8989")
	if err != nil{
		fmt.Println("listening port failed! err: ", err)
		return
	}
	for{
		conn,err :=listen.Accept()//建立连接
		if err != nil{
			fmt.Println("accept failed! err: ", err)
			return
		}
		defer conn.Close()//关闭连接
		newReader := bufio.NewReader(conn)
		txt:=make([]byte,128)
		content :=make([]byte,128)
		//循环接收信息并处理
		for{
			n,err := newReader.Read(txt)
			if err == io.EOF {
				fmt.Println("read file finished!")
				break
			}
			if err != nil{
				fmt.Println("read failed! err: ", err)
				return
			}
			content = append(content, txt[:n]...)
		}
	}
}