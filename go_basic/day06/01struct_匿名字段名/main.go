package main

import "fmt"

//结构体的匿名字段名

//正常结构体
type cat struct{
	name string
	age int
}

//匿名字段结构体   匿名字段结构体只能通过.字段类型调用，结构体中不能有重复类型的字段
type dog struct{
	string
	int
}

func (c cat) speak1() {
	fmt.Print("%v的叫声是这样的,喵喵喵！",c.name)
}

func (d dog) speak2() {
	fmt.Printf("%v的叫声是这样的，汪汪汪！",d.string)
}

func main(){

	//正常结构体实例化
	var c1 = cat{
		"john",
		18,
	}
	c1.speak1()
	fmt.Println(c1.name)
	fmt.Println(c1.age)

	//匿名字段结构体实例化
	var d1 = dog{
		"狗",
		3,
	}
	d1.speak2()
	fmt.Println(d1.string)
	fmt.Println(d1.int)
}