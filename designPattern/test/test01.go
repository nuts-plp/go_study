package main

import (
	"container/ring"
	"fmt"
	"os"
)

type st struct {
}

func (s *st) Pop() interface{} {
	return nil
}
func (s *st) Push() interface{} {
	return nil
}
func main() {
	r := ring.New(8)
	fmt.Println(r.Len())
	fmt.Println(os.Getpagesize())
	//environ := os.Environ()
	file, _ := os.Create("./uxic.txt")
	stat, err := file.Stat()
	fmt.Println(stat, err)
	fmt.Println(file.Name())
	//for _, v := range environ {
	//	fmt.Println(v)
	//}
	fmt.Println(os.Getenv("gopath"))
}
