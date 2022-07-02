package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

//work_pool


type producer struct {
	value int64
}

type consumer struct {
	job *producer
	result int64
}

var produChan =make(chan *producer,100)
var consumeChan =make(chan *consumer,100)
var wg sync.WaitGroup
//产生随机数，存入通道
func produce(proChan chan<- *producer){
	defer wg.Done()
	for j := 0 ; j< 1000;j++{
		x := rand.Int63()
		pro := &producer{
			value:x,
		}
		proChan <- pro
		time.Sleep(time.Second)
	}
	close(produChan)
}


//从通道中取得数据，计算，并放入到consumeChan中
func consume(proChan <-chan *producer, result chan<- *consumer){
	defer wg.Done()
	
	for{//计算每个随机数的和，把和发送给从consumeChan
		job :=<- proChan
		sum:=int64(0)
		n := job.value
		for n > 0{
			sum += n % 10
			n = n / 10
		}
		newResult := &consumer{
			job:job,
			result:sum,
		}
		consumeChan <- newResult

	}

}



func main(){
	wg.Add(25)
	//创建一个goroutine线程产生随机数
	go produce(produChan)

	//for循环创建goroutine池
	for i:=0; i<25;i++{
		go consume(produChan,consumeChan)
	}

	//打印输出consumerChan的值
	for i := range consumeChan {
		fmt.Printf("value:%v、sum:%v\n",i.job.value,i.result)
	}
	wg.Wait()
	

}