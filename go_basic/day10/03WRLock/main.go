package main

import "fmt"
import "sync"

var(
	//读写锁   当有一个线程读时，其他线程也可以读，当有线程写时，其他线程无论读与写都要等待

	//分别比较加锁时与不加锁的数据
	RWlock sync.RWMutex
	wg sync.WaitGroup
	x int = 0
)


//分别定义一个读方法和写方法

func write(){
	defer wg.Done()
	//加入写锁
	RWlock.Lock()
	for i:=0; i<10000; i++ {
		x+=1
	}

	//解除写锁
	RWlock.Unlock()

}

func read(){
	defer wg.Done()

	//加入读锁
	RWlock.RLock()
	fmt.Println(x)

	//解除读锁
	RWlock.RUnlock()

}

func main(){
	wg.Add(50-1)

	go write()
	
	for i := 0; i < 48; i++ {
		go read()
	}

	// go write()
	wg.Wait()

}