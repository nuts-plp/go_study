package scan

import (
	"fmt"
)

//Scan 测试从标准输入
func Scan() {
	var name string
	var age int
	fmt.Scan(&name, &age)
	fmt.Println(name, age)
}

type person struct {
	name string
	age  int
	sex  string
}

//Scanf 测试Scanf
func Scanf() {
	var person1 person
	fmt.Scanf("姓名:%s   age:%d   sex:%s", &person1.name, &person1.age, &person1.sex)
	fmt.Printf("%v", person1)
}

//Scanln 测试Scanln
func Scanln() {
	var (
		name     string
		age      int
		province string
	)
	fmt.Scanln(&name, &age, &province)
	fmt.Println(name, age, province)
}
