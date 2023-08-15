package main

import (
	"fmt"
	"time"
)

var ch chan string

func test() {
	select {
	case a := <-ch:
		fmt.Println(a)
	case <-time.After(time.Second * 3):
		fmt.Println("time over!")

	}
}
func main() {
	fmt.Println("start game!")
	go test()
	ch <- "潘丽萍！ 我有点想念你"
	time.Sleep(time.Second * 4)

}
