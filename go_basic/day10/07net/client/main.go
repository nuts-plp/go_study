package main

//通信客户端
import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main(){
	//1、连接端口

	conn, err := net.Dial("tcp","localhost:20000")
	if err != nil {
		fmt.Println("Client connecting failed! err:",err)
		return
	}
	defer conn.Close()//关闭连接

	inputReader := bufio.NewReader(conn)
	for {
		input,_ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input,"\r\n")
		if strings.ToUpper(inputInfo) == "Q"{
			return
		}
		_,err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("writting to sever failed! err:",err)
			return						
		}
		buf := [512]byte{}
		n,err := conn.Read(buf[:])
		if err != nil{
			fmt.Println("receive message failed! err:",err)
			return
		}
		fmt.Println(string(buf[:n]))

	}
	//2、建立连接
	//3、发送信息


}