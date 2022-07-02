package main
import(
	"net"
	"fmt"
)
/*
	1、建立连接
	2、发送信息
	3、关闭连接
*/

func main(){
	//建立连接
	
	conn,err :=net.Dial("tcp","127.0.0.1:8989")
	if err != nil {
		fmt.Println("Error connecting! err:",err)
		return
	}
	defer conn.Close()//关闭链接
	var txt string
	//循环发送信息
	for{
		fmt.Print("请输入要发送的信息:")
		fmt.Scanln(&txt)
		_,err := conn.Write([]byte(txt))
		if err != nil {
			fmt.Println("message senting failed! err:",err)
			return
		}

	}
}