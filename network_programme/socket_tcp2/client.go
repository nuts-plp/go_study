package main

import (
	"fmt"
	"net"
)

func main() {
	dial, _ := net.Dial("tcp", "127.0.0.1:8888")

	defer dial.Close()

	for {
		var input string
		fmt.Scanln(&input)

		dial.Write([]byte(input))
		buf := make([]byte, 4096)
		n, _ := dial.Read(buf)
		for n <= 0 {
			fmt.Print(buf[:n])
		}
		fmt.Println()

	}

}
