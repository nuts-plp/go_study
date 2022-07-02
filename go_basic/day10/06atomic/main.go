package main

import (
	"fmt"
	"sync/atomic"
)

//原子操作
var x = int64(0)

// func add(){
// 	x += 1
// 	fmt.Println(x)
// }
func main(){
	//
	// for i:=0; i<10000; i++ {
	// 	go add()
	// }
	for i :=0; i<10000; i++ {
		go func (){
			 x :=atomic.AddInt64(&x,1)
			 fmt.Println(x)
		}()
	}

}