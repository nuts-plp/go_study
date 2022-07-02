package main


//当两个线程操控同一个变量时，会因为没有交互而同时对变量操作，结果变量只进行一次
//可以通过lock锁控制同一时间只有一个线程操作变量





import (
	"fmt"
	"sync"
)
var(
	lock sync.Mutex
	wg sync.WaitGroup
	x int = 0
)


//自增运算
func f1(){
	defer wg.Done()
	lock.Lock()
	for i:=0;i<100000;i++ {
		x+=1
	}
	lock.Unlock()
}


func main(){
	//lock锁
	wg.Add(2)
	go f1()

	go f1()
	wg.Wait()
	fmt.Println(x)
}