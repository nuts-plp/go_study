package main

import (
	"fmt"
	"sync"
)

func main() {
	var b = make(chan int, 100)
	for i := 0; i < 100; i++ {
		b <- i
	}
	v := &sync.Pool{
		//New: func() interface{} {
		//	return 0
		//},
	}
	//c := sync.Cond{}
	v.Put(190)
	fmt.Println(v.Get())
	fmt.Println(v.Get())
	//group := sync.WaitGroup{}
	//for i := 0; i < 100; i++ {
	//	group.Add(1)
	//	go func() {
	//		a, ok := <-b
	//		if !ok {
	//			return
	//		}
	//		fmt.Println(a)
	//		group.Done()
	//	}()
	//}
	//
	//group.Wait()
	//var c sync.WaitGroup
}
