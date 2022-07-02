package main

import "fmt"

func main(){
	//panic 和 recover
	f1()
	// f2()
}

//panic 和 recover
func f1(){
	defer func() {
		er := recover()
		fmt.Println("死了都要爱。。。")
		fmt.Println(er)
	}()

	panic("坏了！我们被发现了！")//执行到这里程序就要退出

	fmt.Println("this is real love")//永远无法执行的代码

}

func f2(){

	panic ("wow ! 我们被发现了！")

	fmt.Println("代码永远执行不到的地方")
}