package main

import (
	"fmt"

	"sync"
)

var wg sync.WaitGroup
var b =make(chan int,6)
func main() {
	// //通道初识
	/*
	对于有缓冲通道，存入数据即使不取也没有问题，但是对于无缓冲通道，存入数据必须取而且是全取，否则造成死锁
	如果取得的数据超出存入的数量，为通道类型的默认值

	对于无论是否有缓冲区的通道而言，如果没有向通道传值并且没有显式地关闭通道，就会导致死锁。如果有显式
	的关闭通道，无论是否传值都不会导致死锁

	往通道存入超出容量的数据也会导致死锁
	*/

	// wg.Add(2)
	// go func() {
	// 	for i:=0; i<100; i++ {
	// 		b <-i
	// 		fmt.Println("hello! 潘丽萍",i)
	// 		fmt.Println("通道b发送了！")
	// 	}
	// 	close(b)
	// 	wg.Done()
	// }()

	// go func(){	
	// 	for {
	// 		x ,ok:= <- b
	// 		if !ok{
	// 			break
	// 		}
	// 		fmt.Println(x)
	// 	}
	// 	wg.Done()
	// }()
	// wg.Wait()


	wg.Add(2)
	go func() { 
		//给通道b传值
		b<- 10
		b<- 100
		b<- 1000
		b<- 10000
		b<- 1000000
		
		// for i := 0; i < 100000; i++ {
		// 	b<- i
		// }
		
		defer wg.Done()
		defer close(b)
	}()

	go func() {
		//从通道b取值
		<- b
		<- b
		<- b //以上这些取值但并没有操作
		x ,ok:= <- b//取值，把取得的值赋值给x
		fmt.Println(x,ok)

		wg.Done()
		fmt.Println()
	}()

	wg.Wait()

}