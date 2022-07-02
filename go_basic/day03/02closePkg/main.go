package main

import "fmt"

func main(){
	
	//闭包
	//把fun2作为参数传给fun3，返回得到一个无参函数
	v :=fun3(fun2)

	//然后把返回得到的无参函数传给fun1
	fun1(v)

}

//闭包
//定义一个函数fun3对f2进行封装  即把fun2作为参数传进fun3
//把这个函数fun3作为参数传给fun1
func fun3(ff func(int,int))func(){
	temp := func(){
		ff(5,6)
		fmt.Println("this is fun3. what happing")
	}

	return temp
}
//无参函数作为参数
func fun1(f func()){
	fmt.Println("this is fun1!")
	f()
}

//两个int数
func fun2(x,y int){
	fmt.Println("this is fun2!")
	fmt.Println(x + y)
}