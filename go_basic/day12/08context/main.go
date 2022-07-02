package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var flag bool
var ch = make(chan bool,3)

//用全局变量flag退出死循环
func f1(){
	defer wg.Done()
Loop:
	for{
		fmt.Println("樊雪姨！")
		time.Sleep(time.Millisecond*500)
		
		if flag{
			break Loop
		}
	}

}
//用管道退出死循环
func f2(){
Loop:
	for{
		fmt.Println("周小林")
		time.Sleep(time.Millisecond*500)
		select{
		case <- ch:
			break Loop

		default:

		}
	}
}
//用context退出死循环
func f3(ctx context.Context){
	defer wg.Done()
Loop:
	for{
		fmt.Println("潘丽萍！")
		time.Sleep(time.Millisecond*500)
		
		select{
			
		case <- ctx.Done():
			break Loop
		default:

		}
	}
}

func main(){
	wg.Add(1)
	//借助全局变量退出死循环
	// go f1()
	
	//借助管道退出死循环
	// go f2()
	ctx, cancel := context.WithCancel(context.Background())
	//借助context退出死循环
	go f3(ctx)
	time.Sleep(time.Second*5)

	// flag = true
	// ch<- true


	cancel()
	wg.Wait()
}