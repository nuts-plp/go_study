package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	share := false
	var wg sync.WaitGroup
	var l sync.Mutex
	var c = sync.NewCond(&l)
	wg.Add(2)
	go func() {
		c.L.Lock()
		for share == false {
			fmt.Println("goroutine1 waiting.....")
			c.Wait()
		}
		fmt.Println("goroutine1", share)
		c.L.Unlock()
		wg.Done()

	}()

	go func() {
		c.L.Lock()
		//for share == false {
		//	fmt.Println("goroutine2 waiting.....")
		//	c.Wait()
		//}
		fmt.Println("goroutine2 waiting.....")
		c.Wait()
		fmt.Println("goroutine2", share)
		c.L.Unlock()
		wg.Done()

	}()
	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	share = true
	//c.Broadcast()
	c.Signal()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()

	time.Sleep(time.Second * 5)
	wg.Wait()
}
