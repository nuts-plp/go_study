package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ch = make(chan int, 20)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		ch <- i
		wg.Add(1)
	}
	for i := 0; i < 20; i++ {
		go func(n int) {
			select {
			case v := <-ch:
				time.Sleep(time.Duration(v) * time.Microsecond)
				fmt.Println(n, "->", v)
			case <-time.After(time.Microsecond):
				fmt.Println("time out !!!")
			default:
				fmt.Println("default setting!!")

			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}
