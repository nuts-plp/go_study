package main

import "fmt"

func main() {
	buf := []byte("hjassd")
	fmt.Println(string(buf), "----||----", string(buf[:len(buf)]))

}
