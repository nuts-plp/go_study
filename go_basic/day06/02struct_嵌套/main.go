package main

import "fmt"

//结构体的嵌套    就是匿名结构体的另一种使用

type animal struct{
	name string
	age int

}

type dog struct{
	gender string
	animal //以animal作为变量  匿名结构体

}

type cat struct{
	size string
	animal animal
}

func main(){
	var d1 = dog{
		gender:"公",
		animal:animal{
			name:"peter",
			age:5,
		},
	}
	fmt.Println(d1.gender)

	//下面两种方式访问结果相同
	fmt.Println(d1.name)
	fmt.Println(d1.animal.name)

	var c1 = cat{
		size : "little",
		animal :animal {
			name:"坚果",
			age:2,
		},
	}
	fmt.Println(c1.size)
	// fmt.Println(c1.name)
	// fmt.Println(c1.an.name)//通过级层访问的方式搜索成员变量



}