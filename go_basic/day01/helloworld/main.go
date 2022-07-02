package main	//声明main包  表明当前是一个可执行文件

import "fmt"//导入内置的fmt包

//go语言不支持函数外写语句,但可以声明变量


//变量的声明为静态，必须指定变量的类型，   同时go语言声明变量之后必须使用（减小可执行文件的大小），否则报错
//变量的命名推荐使用驼峰命名法

//全局变量声明
//只声明变量
var name string//字符串类型的声明
//直接声明且赋值
var age int//int类型的声明
var isOk bool//布尔类型的声明

//批量声明变量
// var(
// 	name string//字符串变量默认为""
// 	age int//int默认为0
// 	isOk bool//布尔默认为false
// )

func main() {//main函数 程序执行入口
	//全局变量的赋值
	name = "潘丽萍"
	age = 19
	isOk = true
	//类推变量的声明,根据数据推算数据类型
	var s = 19
	fmt.Println(s)
	
	//简短变量声明（只能在函数体内使用，局部变量）常用的方式
	v :="你好"//它会根据数据来判定类型
	fmt.Println(v)

	fmt.Print(name)
	fmt.Printf("、年龄：%d",age)//%s占位符，格式化打印
	fmt.Println(isOk)//打印后会自动在末尾加上换行符\n 自动换行
	fmt.Println("你好！潘丽萍")
	
}
