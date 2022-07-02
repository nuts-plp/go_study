package main

import (
	"fmt"
	"sync"
	"time"
	// "math/rand"
)

//使用waitGroup等待所有的goroutine结束

var wg sync.WaitGroup

func f1(i int){
	defer time.Sleep(time.Second * 3)
	fmt.Println("你好啊  潘丽萍！",i)
	defer wg.Done()
}

func main(){
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()

}