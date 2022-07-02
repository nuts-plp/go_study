package main

import "fmt"

//type关键词

	
	

func main(){
	//type关键词

	//1、自定义类型
	type myType rune//将myType定义为rune类型

	//用自定义类型声明一个变量
	var a [] myType 


	//自定义类型的变量初始化
	a = make([]myType,9)
	fmt.Printf("%T\n",a)

	//2、给类型起别名
	type yourType = int

	var b yourType

	//给起别名的变量初始化
	b = int(9)
	fmt.Printf("%T",b)
	

}