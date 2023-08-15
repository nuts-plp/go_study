package main

import (
	"fmt"
	"runtime"
)

//崩溃时传递上下文信息
type printContext struct {
	str string
}

//保护方式允许一个函数

func protectFunc(ent func()) {
	//延迟处理
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("运行时错误")
		default:
			fmt.Println("其他错误")
		}
	}()
	ent()
}

func main() {
	fmt.Println("运行之前！")
	protectFunc(func() {
		fmt.Println("手动宕机之前")
		panic(&printContext{"手动触发宕机"})
		fmt.Println("手动宕机之后")
	})
	fmt.Println("被动宕机之前")
	var p *int
	*p = 9
	fmt.Println("被动宕机之后")
	fmt.Println("运行之后")
}
