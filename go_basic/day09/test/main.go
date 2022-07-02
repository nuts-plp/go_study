package main
import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

//工作内容
type job struct {
	num int64
}

//工作结果
type result struct {
	job *job
	sum int64
}

//实例化两个通道，一个用作接收随机数，另一个用作接受结果
var proChan = make(chan *job,100)
var resuChan = make(chan *result,100)

//创建一个waiteGroup
var wg sync.WaitGroup

//创建一个方法，用来产生随机数
func jobMaker(proChan chan<- *job){
	defer wg.Done()
	defer close(proChan)

	//种下一个时间种子
	rand.Seed(time.Now().UnixMicro())
	for i:=0;i<100;i++{
		r :=rand.Int63()
		//把结构体指针存入通道
		newJob := &job{
			num:r,
		}
		proChan <- newJob
		//每产生一个就休息一会儿
		time.Sleep(time.Second)
	}
	
}

//此方法用于接收proChan通道中的数据，并且计算，将结果重新发送到resuChan通道中
func worker(proChan <-chan *job,resuChan chan<- *result){
	defer wg.Done()
	// defer close(resuChan)

	//遍历通道中得值并计算，把计算结果发送到resuChan
	
		receiver := <-proChan
		n := receiver.num
	
		sum := int64(0)

		//循环计算随机数的和
		for {
			
			sum = n % 10
			n = n / 10
			if n < 0 {
				break
			}

		}
		newResult := &result{
			job:receiver,
			sum:sum,
		}

		resuChan<- newResult
	
}

func main(){
	//添加100个计数器
	wg.Add(100)
	//一个goroutine用于产生随机数
	go jobMaker(proChan)
	//99个goroutine用于计算
	for i:=0; i<99;i++{
		go worker(proChan,resuChan)
	}
	for result := range resuChan{
		fmt.Printf("value:%v、sum:%v\n",result.job.num,result.sum)
	}
	wg.Wait()
}