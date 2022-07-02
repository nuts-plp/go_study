package main

import "fmt"

//继承的结构体式实现
type animal struct{
	name string
}

type cat struct{
	age int
	*animal//指针类型
}

func (a animal)f1(){
	fmt.Printf("我是%v,我是凶猛的野兽~\n",a.name)
}

func (c *cat) f2(){
	fmt.Printf("我是%v,我喜欢 喵喵喵~\n",c.name)
}

func main() {
	var a = animal{
		"野兽",
	}
	a.f1()
	fmt.Println(a)


	var c = cat{
		age:1,
		animal:&animal{
			name:"小喵咪",
		},
	}
	
	c.f1()//调用父类方法
	c.f2()//调用自己的方法
	fmt.Println(c)
	fmt.Println(c.name)
	fmt.Println(c.age)

}