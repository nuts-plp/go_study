package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func send() {}

// tasklist ,ipcnfig
func handle() {
	//处理信号
	sigRecvl := make(chan os.Signal, 1)
	sig1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Println("sig:", sig1)

	signal.Notify(sigRecvl, sig1...)

	sigRecv2 := make(chan os.Signal, 1)
	sig2 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Println("sig2:", sig2)
	signal.Notify(sigRecv2, sig2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecvl {
			fmt.Println("recv:1", sig)
		}
		fmt.Println("recv1:", "over")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Println("recv2:", sig)
		}
		fmt.Println("recv2:", "over")
		wg.Done()
	}()
}

func main() {

}
