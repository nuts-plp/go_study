package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//go语言内置的map不是并发安全的，sync.Map是并发安全的 他提供了Store()、Load()、Store()、Delete()、LoadOrStore()、Range()

//go内置的map
var m = make(map[string]int)
var lock sync.Mutex
//sync 包中安全的map
var ms  = sync.Map{}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

//f1方法 对go内置的map操作
func f1(){
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			lock.Lock()
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//f2方法 对sync.Map线程安全的方法进行操作
func f2(){
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			ms.Store(key, n)         // 存储key-value
			value, _ := ms.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func main() {
	f1()
	f2()
	time.Sleep(time.Second)
}