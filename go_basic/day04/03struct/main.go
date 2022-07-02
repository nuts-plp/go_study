package main 

import "fmt"

//struct 结构体

//定义一个结构体
type person struct{
	name ,city string//同样的字段类型也可以写在一行
	age int
}

func main(){

	//结构体实例化
	//1、实例化方法一
	var p1 = person{
		 "john",
		 "武汉",
		 18,
	}
	fmt.Printf("结构体(p1):%T\n",p1)

	//2、实例化方法二
	var p2  person
	p2.name = "mike"
	p2.city = "上海"
	p2.age = 19
	fmt.Printf("结构体(p2):%v\n",p2)

	//3、实例化方法三
	var p3 = new(person)
	p3.name = "nina"
	p3.age = 20
	fmt.Printf("%v\n",p3)

	//4、实例化方法四
	p4:= person{
		name:"miko",
		city:"广州",
		age:21,
	}
	fmt.Println(&p4)

	//5、初始化方法五
	p5 := person{
		"mnkl",
		"深圳",
		18,
	}
	fmt.Println(&p5)

	//匿名结构体
	var dog struct{
		name string
		age int
	}

	//匿名结构体实例化
	dog.name = "小花"
	dog.age = 3
	fmt.Printf("%v\n",dog)

	//取结构体的地址实例化
	p10 := person{}
	fmt.Printf("%v\n",&p10)
	fmt.Printf("%#v\n",&p10)
	fmt.Printf("%#v\n",&(p10.name))
	fmt.Printf("%#v\n",&(p10.age))
	fmt.Println("---------------------------------")

	//利用函数实现实例化结构体
	a := newPerson("潘丽萍","贵阳",19)
	fmt.Println(*a)
	fmt.Println(a)
	


}



//结构体的构造函数
func newPerson(n,c string,a int)*person{

	return &person{
		name:n,
		city:c,
		age:a,
	}
}