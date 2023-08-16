package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if nil != err {
		log.Panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if nil != err {
		log.Panic(err)
	}
	defer conn.Close()
	go func() {
		for {
			buffer := make([]byte, 4096)
			n, udpAddr, _ := conn.ReadFromUDP(buffer)
			fmt.Println(string(buffer))
			command := exec.Command(string(buffer[:n]))
			output, _ := command.CombinedOutput()

			conn.WriteToUDP(output, udpAddr)
		}

	}()

}
