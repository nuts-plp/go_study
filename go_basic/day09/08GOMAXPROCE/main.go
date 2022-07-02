package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCE

var wg sync.WaitGroup
func main(){

	//设置最大线程数
	runtime.GOMAXPROCS(5)

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func(i int){
			defer wg.Done()
			fmt.Println("潘丽萍，我想你了！")
		}(i)
	}
	for i := 0; i < 50; i++ {
		go func(i int){
			defer wg.Done()
			fmt.Println("周小林，我想你了！")
		}(i)
	}
	wg.Wait()
}