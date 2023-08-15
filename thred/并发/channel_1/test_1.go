package main

import (
	"fmt"
	"time"
)

var cha chan string

func test() {
	select {
	case ch, ok := <-cha:
		if ok {
			fmt.Printf("%v", ch)
		}
	case <-time.After(time.Second * 2):
	}
}
func main() {
	go test()
	fmt.Print("ghajhdsgf")
	cha <- "你好！潘丽萍"
	time.Sleep(time.Second * 4)
}
