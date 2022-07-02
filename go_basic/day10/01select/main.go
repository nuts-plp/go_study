package main
import "fmt"

func main(){
	//select
	//在select语句中如果条件都满足，那就随机运行一个case、如果有不满足的那就先运行满足的那一个

	ch := make(chan int,10)
	for i := 0; i < 10; i++ {
		select{
		
		case x:= <- ch:
			fmt.Println(x)

		case ch <- i:

		// default:
			// fmt.Println(i)
		}
	}
}