package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var pool sync.Pool

//func testTCP(){
//	listen, err := net.Listen("tcp4", "localhost:8080")
//	if err != nil {
//		return
//	}
//	for {
//		accept, err := listen.Accept()
//		if err != nil {
//			break
//		}
//		v := strings.Join([]string{accept.RemoteAddr().String(), accept.LocalAddr().String()}, "|")
//		fmt.Println(v)
//		accept.Write([]byte(v))
//		fmt.Println(accept.RemoteAddr().String(), accept.LocalAddr().String(), time.Now().UTC())
//		pool.Put(accept)
//	}
//}
func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":8081")
	udp, _ := net.ListenUDP("udp", addr)
	defer udp.Close()
	for {
		buf := [40]byte{}
		n, err := udp.Read(buf[:])
		fmt.Println(n, err)
		fmt.Println(buf)
		fmt.Println(udp.LocalAddr().String())
		time.Sleep(time.Second * 5)
	}

	//for {
	//	n, err := udp.Write([]byte("I am num1"))
	//	fmt.Println(n, err)
	//	n, err = udp.WriteTo([]byte("I am num1"), addr)
	//	fmt.Println(n, err)
	//	n, err = udp.WriteToUDP([]byte("I am num1"), addr)
	//	fmt.Println(n, err)
	//	n, n1, err := udp.WriteMsgUDP([]byte("I am num1"), []byte("额外"), addr)
	//	fmt.Println(n, n1, err)
	//
	//	time.Sleep(time.Second * 10)
	//}
}
