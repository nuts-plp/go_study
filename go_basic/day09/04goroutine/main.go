package main

import (
	"fmt"
	"time"
)

//并发
func f1 (i int){
	fmt.Println("hello world!",i)
}

func f2(){
	fmt.Println("你好！潘丽萍")
}
func main(){
	// func (){
	// 	for i:=0; i<100000; i++{
	// 		go f1(i)
	// 	}
	// }()
	go f2()
	

	for i := 0; i < 1000; i++{
		go func(){
			fmt.Println("hello 潘丽萍!",i)
		}()
	}
	time.Sleep(time.Second)

}
