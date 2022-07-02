package main
 
import "fmt"
import "sync"

var (
	wg sync.WaitGroup
	once sync.Once
	ch = make(chan int,10)
)

func f1(i int){
	defer wg.Done()
	if i == 10{
		defer once.Do(func(){close(ch)})//关闭通道导致的死锁
	}
	
	ch<-i
}

func main(){
	wg.Add(11)

	//创建了十个线程，只有一个显式的关闭了
	for i := 0; i < 10; i++ {
		go f1(i)
	}
	

	go func (){
		defer wg.Done()
		for i:= range ch{
			fmt.Printf("value: %d\n", i)
		}
	}()
	wg.Wait()

}