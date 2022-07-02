package main

import "fmt"
func main(){
	//指针
	//取地址&
	//根据地址取值
	a := "永远不要高估自己的能力"
	b := &a		//把a的内存地址赋值给b   b是一个字符串类型指针  *string
	fmt.Printf("%T",b)
	fmt.Println()
	fmt.Println(b)
	fmt.Println(*b)//根据变量b存取的地址取值

	var x *int//变量x为int指针数据类型
	// var x = new(int)//开辟内存空间
	*x = 100//为地址赋值 ，但是在内存中并没有为地址开辟空间，所以会出错
	fmt.Println(*x)
}