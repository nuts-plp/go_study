package main

import "fmt"

//一种特殊的函数，只能被特殊的类型调用
//值接收者和指针接收者  只能被只能被指定类型调用
//不允许有多个接收者

type cat struct{
	name string
	age int
}




//指针接收者
func (c *cat) eat() {
	fmt.Printf("我是%s,喵喵喵\n",c.name)
}

//值接收者
func (c cat) run(){

	fmt.Printf("我是%s,我喜欢到处乱跑\n",c.name)

}

//
func (p person)drink(){

	fmt.Printf("我是%s,我渴了！\n",p.name)
}


func main(){

	a := person{
		"lkio",
		"女",
	}
	b := cat{
		"uidn",
		3,
	}
	a.drink()
	b.run()
	b.eat()
	fmt.Println("-------------------------------")

	w := myInt(16)
	w.ex()
	
}

type myInt  int

func (m myInt) ex(){
	fmt.Println("啊啊啊,毁我三观啊")
}



type person struct {
	name string
	gender string
}