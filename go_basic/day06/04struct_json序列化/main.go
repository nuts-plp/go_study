package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json序列化与反序列化

type Address struct {
	Province string
	Location string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {

	var a = Address{
		"湖北",
		"江夏区武纺",
	}

	//序列化单层结构体   结构体-->json字符串
	data1, err1 := json.Marshal(a)
	if err1 != nil {
		fmt.Println("Marshal failed to parse!")
		return
	}
	fmt.Printf("%v\n", data1)
	fmt.Printf("%#v\n", data1)
	fmt.Println(data1)
	//反序列化     json字符串--->结构体
	var data11 Address
	err11 := json.Unmarshal([]byte(data1), &data11) //用指针是为了能够修改data11的值
	if err11 != nil {
		fmt.Println("Unmarshal failed to parse!")
	}
	fmt.Println(data11)

	var p = Person{
		Name: "潘丽萍",
		Age:  19,
		Address: Address{
			Province: "贵州省",
			Location: "黔东南",
		},
	}
	//序列化多层结构体
	data2, err2 := json.Marshal(p)
	if err2 != nil {
		fmt.Println("Msrshal failed to parse!")
	}

	//反序列化多层结构体
	var data22 Person
	err22 := json.Unmarshal([]byte(data2), &data22) //用指针是为了修改的是data22的值
	if err22 != nil {
		fmt.Println("Unmarshal failed to pares!")
	}
	fmt.Println(data22)
	fmt.Printf("%v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Println(p)

}
