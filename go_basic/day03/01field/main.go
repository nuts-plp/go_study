package main

import "fmt"

//变量的作用域
//全局变量
var a int = 22 



func main(){

	//函数体内的局部变量
	fmt.Println(f1("潘丽萍！"))
	fmt.Println(a)
	//语句块变量
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	
}
func f1(s string)(ss string,i int){
	ss = s + "I miss you"
	m := 90
	return ss,m
}