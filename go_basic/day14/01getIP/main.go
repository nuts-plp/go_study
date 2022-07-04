package main

import (
	"fmt"
	"go_basic/day14/01getIP/getIP"
)

func main() {
	IP, err := getIP.GetOutboundIP()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(IP)
}
