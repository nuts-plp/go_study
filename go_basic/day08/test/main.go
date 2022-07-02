package main
import (
	"fmt"
	"time"
)

func main(){
	//1、获取当前时间，格式化输出
	fmt.Println(time.Now().Format("2006/01/02 03:04:05"))
	now := time.Now().UnixMicro()
	// ticker := time.Tick(time.Second)
	for i := 0; i <100; i++ {
		if i % 17 == 0{
			time.Sleep(time.Second*1)
		}
	}
	later := time.Now().UnixMicro()
	fmt.Println(later-now)
	
}