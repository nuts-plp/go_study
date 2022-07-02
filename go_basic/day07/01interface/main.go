package main

import "fmt"

//接口的定义及实现


type animal interface{
	speak()
}


type dog struct {
	name string
}
//
func(d dog)speak(){
	fmt.Printf("我是一只狗，我的名字是:%s\n",d.name)
}

type cat struct{
	name string
}

//cat的接收者
func (c cat)speak(){
	fmt.Printf("我是一只猫，我叫：%s\n",c.name)
}

//声明一个speaker
func speaker(a animal){
	a.speak()
}

func main(){
	d := dog{
		"花花",
	}
	c := cat{
		"琪琪",
	}
	speaker(d)
	speaker(c)

}
//一个特殊的方法  无参无返回值，类似于java的静态代码块
//在main方法执行前执行
func init(){
	fmt.Println("我是init()方法，我在main()方法执行之前执行！")

}